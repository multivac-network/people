package data

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type PersonDataStore struct {
	GraphDataStore
}

type PersonUpdate struct {
	Current map[string]interface{}
	Previous map[string]interface{}
}

var s = &PersonDataStore{}

func Initialize(uri, username, password string) {
	s.initialize(uri, username, password)
}

func Store() *PersonDataStore{
	return s
}

func (store *PersonDataStore) FindAll() ([]*Person, error){
	items, err := store.read("MATCH (p:Person) RETURN p.FirstName, p.LastName, p.id, p.Title", nil)
	if err != nil {
		panic(err)
	}
	result := make([]*Person, 0)
	records :=items.([]*db.Record)
	for _, v := range records {
		id, _ := v.Get("p.id")
		firstName, _ := v.Get("p.FirstName")
		lastName, _ := v.Get("p.LastName")
		var title string
		if t, _ := v.Get("p.Title"); t != nil{ title = t.(string) }
		result = append(result, &Person{
			Id:id.(string),
			FirstName:firstName.(string),
			LastName:lastName.(string),
			Title: &title,
		})
	}
	return result, nil
}

func (store *PersonDataStore) FindByOrganizationId(id string) ([]*Person, error){
	println("OrganizationID: " + id)
	items, err := store.read("MATCH (p:Person)-[:RESOURCE_OF]->(Organization{id:$id}) RETURN p.id, p.FirstName, p.LastName, p.Title",
		map[string]interface{}{"id": id})
	if err != nil {
		panic(err)
	}
	result := make([]*Person, 0)
	records :=items.([]*db.Record)
	for _, v := range records {
		id, _ := v.Get("p.id")
		firstName, _ := v.Get("p.FirstName")
		lastName, _ := v.Get("p.LastName")
		var title string
		if t, _ := v.Get("p.Title"); t != nil{ title = t.(string) }
		result = append(result, &Person{
			Id:id.(string),
			FirstName:firstName.(string),
			LastName:lastName.(string),
			Title: &title,

		})
	}
	return result, nil
}

func (store *PersonDataStore) FindById(id string) ([]*Person, error){
	items, err := store.read("MATCH (p:Person{id: $id}) RETURN p.FirstName, p.LastName, p.id, p.Title",
		map[string]interface{}{"id": id})
	if err != nil {
		panic(err)
	}
	result := make([]*Person, 0)
	records :=items.([]*db.Record)
	for _, v := range records {
		id, _ := v.Get("p.id")
		firstName, _ := v.Get("p.FirstName")
		lastName, _ := v.Get("p.LastName")
		var title string
		if t, _ := v.Get("p.Title"); t != nil{ title = t.(string) }
		result = append(result, &Person{
			Id:id.(string),
			FirstName:firstName.(string),
			LastName:lastName.(string),
			Title: &title,
		})
	}
	return result, nil
}

// Creates a new person with a auto-generated UUIDv4
func (store *PersonDataStore) Create(person Person) (Person, error) {
	params := map[string]interface{}{"FirstName": person.FirstName, "LastName": person.LastName, "name": person.FirstName + " " + person.LastName}
	if person.Title != nil {
		params["Title"] = *person.Title
	}
	p, err := store.write(
		"CREATE (p:Person) " +
			"SET p.id = $id, " +
			"p.name = $name, " +
			"p.FirstName = $FirstName, " +
			"p.LastName = $LastName, " +
			"p.Title = $Title " +
			"RETURN p",
			params,
		)

	if p != nil {
		records := p.([]*db.Record)
		object, _ := records[0].Get("p")
		person := object.(dbtype.Node)
		id, _ := person.Props["id"]
		firstName, _ := person.Props["FirstName"]
		lastName, _ := person.Props["LastName"]
		var title string
		if t, _ := person.Props["Title"]; t != nil{ title = t.(string) }
		return Person{
			Id:id.(string),
			FirstName:firstName.(string),
			LastName:lastName.(string),
			Title: &title,
		}, err
	}

	return person, err
}

func (store *PersonDataStore) Delete(person Person) (Person, error) {
	return Person{}, nil
}

func (store *PersonDataStore) Update(person Person) (*PersonUpdate, error) {
	params := map[string]interface{}{"id": person.Id, "FirstName": person.FirstName, "LastName": person.LastName, "name": person.FirstName + " " + person.LastName}
	if person.Title != nil {
		params["Title"] = *person.Title
	}
	out, err := store.write(
		"MATCH (previous:Person{id: $id}) " +
			"MERGE (current:Person{id: $id}) " +
			"ON MATCH " +
			"SET current.name = $name, " +
			"current.FirstName = $FirstName, " +
			"current.LastName = $LastName, " +
			"current.Title = $Title " +
			"RETURN current, previous",
		params,
	)
	record := out.([]*db.Record)[0]
	current, _ := record.Get("current")
	previous, _ := record.Get("previous")

	return &PersonUpdate{Current: current.(dbtype.Node).Props, Previous: previous.(dbtype.Node).Props}, err
}
