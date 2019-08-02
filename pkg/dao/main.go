package database

type Database struct {
	Redis
	MongoDB
}

func (database *Database) InitDatabase() {
	database.Redis.InitRedis()
	database.MongoDB.InitMongoDB()
}
