package service

import (
	"fmt"
)

type Test2Service struct{}


func (*Test2Service) Ping () string {
	fmt.Println(&dao)
	return "test2 service pong"
}
