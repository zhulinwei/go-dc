package config

import (
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util"
	"log"
	"sync"
)

const (
	externalConfigPath = "/var/config.yaml"
	defaultConfigPath  = "./configs/config.yaml"
)

var serverConfig model.ServerConfig
var serverConfigMutex sync.Mutex
var isServerConfigParseFinish bool

func parseServerConfig(){
	serverConfigMutex.Lock()
	defer serverConfigMutex.Unlock()

	var serverConfigPath string
	// 优先使用外部配置文件，后使用默认配置文件（用于容器化方案）
	if exist, _ := util.GetHelper().IsPathExists(externalConfigPath); exist {
		serverConfigPath = externalConfigPath
	} else {
		serverConfigPath = defaultConfigPath
	}

	var err error
	if serverConfig, err = util.GetHelper().ParseServerConfig(serverConfigPath); err != nil {
		log.Fatalf("parse config fail: %v", err)
	}
	isServerConfigParseFinish = true
}

func ServerConfig() model.ServerConfig {
	if !isServerConfigParseFinish {
		parseServerConfig()
	}
	return serverConfig
}
