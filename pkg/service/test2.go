package service



type Test2Service struct{}



func (*Test2Service) Ping () string {
	//redisResult := repository.GetReis().Get("xxxx")
	//fmt.Println(redisResult)
	//
	//
	//user := new(model.Test2Model)
	////repository.GetMongoDB().Database("cms").Collection("test").FindOne(context.TODO(), bson.M{"name": "tony"}).Decode(user)
	//
	//repository.GetMongoDB().Database("cms").Collection()
	//fmt.Println(user)
	//mongoResult := &model.Test2Model{}
	//err := repository.GetMongoDB().Database("cms").Collection("test").FindOne(context.Background(), model.Test2Model{}).Decode(mongoResult).Decode()
	//fmt.Println(err)
	//repository.GetMongoDB().Database("cms").Collection("test").Find(context.TODO(), bson.D{{"appId", appId})
	return "test2 service pong"
}
