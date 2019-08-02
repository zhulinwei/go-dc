package service

type Test2Service struct{}


func (*Test2Service) Ping () string {
	return "test2 service pong"
}
