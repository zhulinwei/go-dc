package util

import (
	"context"
	"time"
)

type Helper struct{}

var helper = new(Helper)

func GetHelper () *Helper {
	return helper
}

func (*Helper) GetContent() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}
