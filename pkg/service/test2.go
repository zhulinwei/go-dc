package service



type Test2Service struct{}



func (*Test2Service) Ping () string {
	//redisResult := database.GetReis().Get("xxxx")
	//fmt.Println(redisResult)
	//
	//
	//user := new(model.Test2Model)
	////database.GetMongoDB().Database("cms").Collection("test").FindOne(context.TODO(), bson.M{"name": "tony"}).Decode(user)
	//
	//database.GetMongoDB().Database("cms").Collection()
	//fmt.Println(user)
	//mongoResult := &model.Test2Model{}
	//err := database.GetMongoDB().Database("cms").Collection("test").FindOne(context.Background(), model.Test2Model{}).Decode(mongoResult).Decode()
	//fmt.Println(err)
	//database.GetMongoDB().Database("cms").Collection("test").Find(context.TODO(), bson.D{{"appId", appId})
	return "test2 service pong"
}
