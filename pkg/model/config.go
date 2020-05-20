package model

type ServerConfig struct {
	Log      LogConfig     `yaml:"log"`
	Name     string        `yaml:"name"`
	Mode     string        `yaml:"mode"`
	Grpc     []GrpcConfig  `yaml:"grpc"`
	MySQL    []MySQLConfig `yaml:"mysql"`
	Redis    []RedisConfig `yaml:"redis"`
	MongoDB  []MongoConfig `yaml:"mongo"`
	HttpPort string        `yaml:"httpPort"`
	GrpcPort string        `yaml:"grpcPort"`
}

type LogConfig struct {
	Level         int  `yaml:"level"`
	DisableCaller bool `yaml:"disableCaller"`
}

type GrpcConfig struct {
	Name string `yaml:"name"`
	Addr string `yaml:"addr"`
}

type MySQLConfig struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Addr string `yaml:"addr"`
}

type RedisConfig struct {
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
