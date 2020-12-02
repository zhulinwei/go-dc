package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhulinwei/go-dc/pkg/config"
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util/log"
)

const (
	defaultDB = "db1"
)

type IMySQL interface {
	DBClient() *sql.DB
}

type MySQL struct {
	once      sync.Once
	configs   []model.MySQLConfig
	ClientMap map[string]*sql.DB
}

var mysql *MySQL

func BuildMySQL() IMySQL {
	if mysql == nil {
		fmt.Println("mysql配置：", config.ServerConfig().MySQL)
		mysql = &MySQL{configs: config.ServerConfig().MySQL}
		mysql.init()
	}
	return mysql
}

func (mysql *MySQL) DBClient() *sql.DB {
	return mysql.ClientMap[defaultDB]
}

func (mysql *MySQL) init() {
	mysql.once.Do(func() {
		mysql.ClientMap = make(map[string]*sql.DB, len(mysql.configs))
		for _, mysqlConfig := range mysql.configs {
			client, err := sql.Open(mysqlConfig.Type, mysqlConfig.Addr)
			if err != nil {
				log.Error("mysql connect fail",
					log.String("db", mysqlConfig.Name),
					log.String("error", err.Error()))
				return
			}
			if err = client.Ping(); err != nil {
				log.Error("mysql ping fail",
					log.String("db", mysqlConfig.Name),
					log.String("error", err.Error()))
				return
			}
			// 保存mysql客户端
			mysql.ClientMap[mysqlConfig.Name] = client
		}
	})
}
