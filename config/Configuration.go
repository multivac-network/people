package config

import "os"

type Configuration struct {
	ServiceName string
	Neo4j       *Neo4jConfiguration
}

func LoadConfiguration() *Configuration {

	result := &Neo4jConfiguration{
		Username: os.Getenv("NEO4J_USER"),
		Password: os.Getenv("NEO4J_PASS"),
		URI:      os.Getenv("NEO4J_URI"),
	}

	return &Configuration{Neo4j: result, ServiceName: os.Getenv("SERVICE_NAME")}
}
