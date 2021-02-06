package data

type PersonDataStore struct {
  DataStore
}

func NewPersonDataStore(uri, username, password string) (*PersonDataStore) {
  datastore := &PersonDataStore{}
  datastore.initialize(uri, username, password)
  return datastore
}

func (ds *PersonDataStore) Save(person Person) (interface{}, error) {
  return ds.execute(
    "CREATE (p:Person) SET p.FirstName = $FirstName, p.LastName = $LastName RETURN id(p)",
    map[string]interface{}{"FirstName": person.FirstName, "LastName": person.LastName})
}