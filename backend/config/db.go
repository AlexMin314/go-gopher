package config

type DatabaseConfig struct {
	Mongo MongoConfig
}

type MongoConfig struct {
	ConnectionUri string
}
