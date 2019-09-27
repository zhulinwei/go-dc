package config

import (
	"github.com/zhulinwei/gin-demo/pkg/model"
	"github.com/zhulinwei/gin-demo/pkg/util"
)

const (
	externalConfigPath = "/var/config.yaml"
	defaultConfigPath  = "./configs/config.yaml"
)

var serverConfig model.ServerConfig

func ServerConfig() model.ServerConfig {
	return serverConfig
}

func init() {
	var serverConfigPath string

	// 优先使用外部配置文件，后使用默认配置文件（用于容器化方案）
	if exist, _ := util.GetHelper().IsPathExists(externalConfigPath); exist {
		serverConfigPath = externalConfigPath
	} else {
		serverConfigPath = defaultConfigPath
	}

	var err error
	if serverConfig, err = util.GetHelper().ParseServerConfig(serverConfigPath); err != nil {
		panic(err.Error())
	}
}
