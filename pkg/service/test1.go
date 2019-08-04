package service

type Test1Service struct{}


func (*Test1Service) Ping () string {
	return "test1 service pong"
}

func (*Test1Service) SaveUser () {

}