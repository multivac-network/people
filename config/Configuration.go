package config

import "os"

type Configuration struct {
	ServiceName     string
	Neo4j           *Neo4jConfiguration
	DevelopmentMode string
}

func LoadConfiguration() *Configuration {

	result := &Neo4jConfiguration{
		Username: os.Getenv("NEO4J_USER"),
		Password: os.Getenv("NEO4J_PASS"),
		URI:      os.Getenv("NEO4J_URI"),
	}

	return &Configuration{
		Neo4j: result,
		DevelopmentMode: os.Getenv("DEVELOPMENT_MODE"),
		ServiceName: os.Getenv("SERVICE_NAME"),
	}
}
