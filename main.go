package main

import "fmt"
import "repath.io/config"
import "repath.io/data"

func main() {
  configuration := config.LoadConfiguration()
  fmt.Printf("initializing %s\n", configuration.ServiceName)
  output, _ := helloWorld(configuration.Neo4j.URI, configuration.Neo4j.Username, configuration.Neo4j.Password)
  print(output)
}

func helloWorld(uri, username, password string) (interface{}, error) {
  datastore := data.NewPersonDataStore(uri, username, password)
  result, _ := datastore.Save(data.Person{FirstName: "Joe", LastName:"Kelly"})
  datastore.Close()

	return result, nil
}