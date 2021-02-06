package data

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type DataStore struct {
  driver neo4j.Driver
  session neo4j.Session
}

func (ds *DataStore) initialize(uri, username, password string) {
  driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
  if err != nil {
    panic(err)
  }
  ds.driver = driver
  session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
  ds.session = session
}

func (ds *DataStore) execute(command string, parameters map[string]interface{}) (interface{}, error) {
  return ds.session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		tresult, err := transaction.Run(command, parameters)
		if err != nil { return nil, err }

		if tresult.Next() {
			return tresult.Record().Values[0], nil
		}
		return nil, tresult.Err()
	})
  
}

func (ds *DataStore) Close() {
  ds.session.Close()
  ds.driver.Close()
}