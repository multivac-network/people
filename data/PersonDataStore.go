package data

type PersonDataStore struct {
	GraphDataStore
}

func NewPersonDataStore(uri, username, password string) *PersonDataStore {
	datastore := &PersonDataStore{}
	datastore.initialize(uri, username, password)
	return datastore
}

// Creates a new person with a auto-generated UUIDv4
func (ds *PersonDataStore) Create(person Person) (Person, error) {
	id, err := ds.execute(
		"CREATE (p:Person) SET p.id = $id, p.name = $name, p.FirstName = $FirstName, p.LastName = $LastName RETURN p.id",
		map[string]interface{}{"FirstName": person.FirstName, "LastName": person.LastName, "name": person.FirstName + " " + person.LastName})
	if err == nil {
		person.Id = id.(string)
	}

	return person, err
}
