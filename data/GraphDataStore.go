package data

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type GraphDataStore struct {
	driver  neo4j.Driver
	session neo4j.Session
}

func (gds *GraphDataStore) initialize(uri, username, password string) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic(err)
	}
	gds.driver = driver
	session, err := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil { panic(err) }
	gds.session = session
}

func (gds *GraphDataStore) execute(command string, parameters map[string]interface{}) (interface{}, error) {
	return gds.session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
    parameters["id"] = uuid.New().String()
		tresult, err := transaction.Run(command, parameters)
		if err != nil {
      panic(err)
		}

		if tresult.Next() {
			return tresult.Record().Values()[0], nil
		}
		return nil, tresult.Err()
	})

}

func (gds *GraphDataStore) Close() {
	gds.session.Close()
	gds.driver.Close()
}
