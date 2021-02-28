package data

type PersonDataStore struct {
	GraphDataStore
}

var s = &PersonDataStore{}

func Initialize(uri, username, password string) {
	s.initialize(uri, username, password)
}

func Store() *PersonDataStore{
	return s
}

func (store *PersonDataStore) FindAll() ([]*Person, error){
	items, err := store.read("MATCH (p:Person) RETURN p.FirstName, p.LastName, p.id", nil)
	if err != nil {
		panic(err)
	}
	result := make([]*Person, 0)
	for _, v := range items {
		id, _ := v.Get("p.id")
		firstName, _ := v.Get("p.FirstName")
		lastName, _ := v.Get("p.LastName")
		result = append(result, &Person{
			Id:id.(string),
			FirstName:firstName.(string),
			LastName:lastName.(string),
		})
	}
	return result, nil
}

func (store *PersonDataStore) FindByOrganizationId(id string) ([]*Person, error){
	println("OrganizationID: " + id)
	items, err := store.read("MATCH (p:Person)-[:MemberOf]->(Organization{id:$id}) RETURN p.id, p.FirstName, p.LastName",
		map[string]interface{}{"id": id})
	if err != nil {
		panic(err)
	}
	result := make([]*Person, 0)
	for _, v := range items {
		id, _ := v.Get("p.id")
		firstName, _ := v.Get("p.FirstName")
		lastName, _ := v.Get("p.LastName")
		result = append(result, &Person{
			Id:id.(string),
			FirstName:firstName.(string),
			LastName:lastName.(string),
		})
	}
	return result, nil
}

func (store *PersonDataStore) FindById(id string) ([]*Person, error){
	items, err := store.read("MATCH (p:Person{id: $id}) RETURN p.FirstName, p.LastName, p.id",
		map[string]interface{}{"id": id})
	if err != nil {
		panic(err)
	}
	result := make([]*Person, 0)
	for _, v := range items {
		id, _ := v.Get("p.id")
		firstName, _ := v.Get("p.FirstName")
		lastName, _ := v.Get("p.LastName")
		result = append(result, &Person{
			Id:id.(string),
			FirstName:firstName.(string),
			LastName:lastName.(string),
		})
	}
	return result, nil
}

// Creates a new person with a auto-generated UUIDv4
func (store *PersonDataStore) Create(person Person) (Person, error) {

	id, err := store.write(
		"CREATE (p:Person) SET p.id = $id, p.name = $name, p.FirstName = $FirstName, p.LastName = $LastName RETURN p.id",
		map[string]interface{}{"FirstName": person.FirstName, "LastName": person.LastName, "name": person.FirstName + " " + person.LastName})
	if err == nil {
		person.Id = id.(string)
	}

	return person, err
}
