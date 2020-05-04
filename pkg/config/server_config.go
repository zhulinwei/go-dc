package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util"
)

const (
	externalConfigPath = "/var/config.yaml"
	// defaultConfigPath  = "./configs/config.yaml"
	defaultConfigPath = "/configs/config.yaml"
)

var serverConfig model.ServerConfig
var serverConfigMutex sync.Mutex
var isServerConfigParseFinish bool

func ServerConfig() model.ServerConfig {
	fmt.Println("come in ServerConfig")
	if !isServerConfigParseFinish {
		parseServerConfig()
	}
	return serverConfig
}

func parseServerConfig() {
	fmt.Println("come in config")
	serverConfigMutex.Lock()
	defer serverConfigMutex.Unlock()

	var serverConfigPath string
	// 优先使用外部配置文件，后使用默认配置文件（用于容器化方案）
	if exist, _ := util.IsPathExists(externalConfigPath); exist {
		serverConfigPath = externalConfigPath
	} else {
		serverConfigPath = defaultConfigPath
	}

	var err error
	if serverConfig, err = util.ParseServerConfig(serverConfigPath); err != nil {
		log.Fatalf("parse config fail: %v", err.Error())
	}
	isServerConfigParseFinish = true
}
