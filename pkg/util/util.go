package util

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/zhulinwei/go-dc/pkg/model"
	"gopkg.in/yaml.v3"
)

func CommonContent() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func IsPathExists(path string) (bool, error) {
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

func ParseServerConfig(filePath string) (model.ServerConfig, error) {
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
