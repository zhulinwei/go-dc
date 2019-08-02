package database

type Database struct {
	Redis
	MongoDB
}

func init() {
	database := new(Database)
	database.Redis.InitRedis()
	database.MongoDB.InitMongoDB()
}
