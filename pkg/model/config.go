package model

type GrpcConfig struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
}

type MySQLConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Addr string `yaml:"addr"`
}

type ReidsConfig struct {
	DB       int    `yaml:"db"`
	Name     string `yaml:"name"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type MongoConfig struct {
	Name     string `yaml:"name"`
	Addr     string `yaml:"addr"`
	Database string `yaml:"database"`
}

type ServerConfig struct {
	Name    string        `yaml:"name"`
	Port    string        `yaml:"port"`
	Grpc    []GrpcConfig  `yaml:"grpc"`
	MySQL   []MySQLConfig `yaml:"mysql"`
	Redis   []ReidsConfig `yaml:"redis"`
	MongoDB []MongoConfig `yaml:"mongo"`
}
