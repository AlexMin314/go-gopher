package config

import "time"

type DatabaseConfig struct {
	Mongo MongoConfig
}

type MongoConfig struct {
	ConnectionUri     string
	ConnectionTimeout time.Duration
	PingTimeout       time.Duration
	DbName            string
	MaxTodoCount      int
}
