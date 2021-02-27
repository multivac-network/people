package data

import (
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	gds.session = session
}

func (gds *GraphDataStore) read(command string, params map[string]interface{}) ([]*neo4j.Record, error) {

	out, err := gds.session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		records := make([]*neo4j.Record, 0)
		tresult, err := transaction.Run(command, params)
		if err != nil {
			panic(err)
		}

		for tresult.Next() {
			records = append(records, tresult.Record())
		}
		return records, tresult.Err()
	})
	if err != nil {
		panic(err)
	}
	return out.([]*neo4j.Record), nil
}

func (gds *GraphDataStore) write(command string, parameters map[string]interface{}) (interface{}, error) {
	return gds.session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		parameters["id"] = uuid.New().String()
		tresult, err := transaction.Run(command, parameters)
		if err != nil {
			panic(err)
		}

		if tresult.Next() {
			return tresult.Record().Values[0], nil
		}
		return nil, tresult.Err()
	})
}

func (gds *GraphDataStore) Close() {
	gds.session.Close()
	gds.driver.Close()
}
