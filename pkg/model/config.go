package model

type ReidsConfig struct {
	DB       int    `yaml:"db"`
	Name     string `yaml:"name"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type MongoConfig struct {
	Name       string `yaml:"name"`
	Addr       string `yaml:"addr"`
	Database   string `yaml:"database"`
}

type Config struct {
	Name  string `yaml:"name"`
	Redis []ReidsConfig
	Mongo []MongoConfig
}
