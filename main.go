package main

import (

	"fmt"
	"log"
	"net/http"
	"repath.io/config"
	"repath.io/data"
)

func main() {
	configuration := config.LoadConfiguration()
	fmt.Printf("starting %s service\n", configuration.ServiceName)
	datastore := data.NewPersonDataStore(configuration.Neo4j.URI, configuration.Neo4j.Username, configuration.Neo4j.Password)
	defer datastore.Close()

	// initialize API and pass datastore
	http.HandleFunc("/echo", echoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		log.Println(err)
	}
}
