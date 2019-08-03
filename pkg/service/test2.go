package service



type Test2Service struct{}



func (*Test2Service) Ping () string {
	//redisResult := dao.GetReis().Get("xxxx")
	//fmt.Println(redisResult)
	//
	//
	//user := new(model.Test2Model)
	////dao.GetMongoDB().Database("cms").Collection("test").FindOne(context.TODO(), bson.M{"name": "tony"}).Decode(user)
	//
	//dao.GetMongoDB().Database("cms").Collection()
	//fmt.Println(user)
	//mongoResult := &model.Test2Model{}
	//err := dao.GetMongoDB().Database("cms").Collection("test").FindOne(context.Background(), model.Test2Model{}).Decode(mongoResult).Decode()
	//fmt.Println(err)
	//dao.GetMongoDB().Database("cms").Collection("test").Find(context.TODO(), bson.D{{"appId", appId})
	return "test2 service pong"
}
