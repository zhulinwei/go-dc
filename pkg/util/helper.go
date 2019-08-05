package util

import (
	"context"
	"time"
)

type Util struct{}

var util = new(Util)

func GetUitl () *Util {
	return util
}

func (*Util) GetContent() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}
