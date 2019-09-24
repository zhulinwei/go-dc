package util

import (
	"context"
	"fmt"
	"github.com/zhulinwei/gin-demo/pkg/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"time"
)

type Helper struct{}

var helper = new(Helper)

func GetHelper() *Helper {
	return helper
}

func (Helper) GetContent() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func (Helper) ParseServerConfig(filePath string) (model.Config, error) {
	var err error
	var fileBytes []byte
	var serverConfig model.Config
	if fileBytes, err = ioutil.ReadFile(filePath); err != nil {
		fmt.Println(err)
		return serverConfig, err
	}
	fmt.Println(fileBytes)
	if err = yaml.Unmarshal(fileBytes, &serverConfig); err != nil {
		return serverConfig, err
	}
	fmt.Println(serverConfig)
	return serverConfig, nil
}