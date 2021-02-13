package main

import (
	"fmt"
	"repath.io/config"
	"repath.io/data"
)

func main() {
	configuration := config.LoadConfiguration()
	fmt.Printf("starting %s service\n", configuration.ServiceName)
	datastore := data.NewPersonDataStore(configuration.Neo4j.URI, configuration.Neo4j.Username, configuration.Neo4j.Password)

	// initialize API and pass datastore

	defer datastore.Close()
}
