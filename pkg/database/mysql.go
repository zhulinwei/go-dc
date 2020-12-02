package database

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zhulinwei/go-dc/pkg/config"
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util/log"
)

type IMySQL interface {
	DB1Client() *sql.DB
}

type MySQL struct {
	once      sync.Once
	configs   []model.MySQLConfig
	ClientMap map[string]*sql.DB
}

var mysql *MySQL

func BuildMySQL() IMySQL {
	if mysql == nil {
		mysql = &MySQL{configs: config.ServerConfig().MySQL}
		mysql.init()
	}
	return mysql
}

func (mysql *MySQL) DB1Client() *sql.DB {
	return mysql.ClientMap["db1"]
}

func (mysql *MySQL) init() {
	mysql.once.Do(func() {
		mysqlClientMap := make(map[string]*sql.DB, len(mysql.configs))
		for _, mysqlConfig := range mysql.configs {
			client, err := sql.Open(mysqlConfig.Type, mysqlConfig.Addr)
			if err != nil {
				log.Error("mysql connect fail", log.String("error", err.Error()))
				return
			}
			if err = client.Ping(); err != nil {
				log.Error("mysql ping fail", log.String("error", err.Error()))
				return
			}
			// 保存mysql客户端
			mysqlClientMap[mysqlConfig.Name] = client
		}
	})
}
