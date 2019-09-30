package util

import (
	"context"
	"fmt"
	"github.com/zhulinwei/go-dc/pkg/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
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

func (Helper) IsPathExists(path string) (bool, error) {
	// 若返回的错误为nil,说明文件或文件夹存在
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	// 若IsNotExist判断为true,说明文件或文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	// 若返回的错误为其它类型,则不确定是否在存在，可视为不存在
	return false, err
}

func (Helper) ParseServerConfig(filePath string) (model.ServerConfig, error) {
	var err error
	var fileBytes []byte
	var serverConfig model.ServerConfig
	if fileBytes, err = ioutil.ReadFile(filePath); err != nil {
		fmt.Println(err)
		return serverConfig, err
	}
	if err = yaml.Unmarshal(fileBytes, &serverConfig); err != nil {
		return serverConfig, err
	}
	return serverConfig, nil
}