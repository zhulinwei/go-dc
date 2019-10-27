package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhulinwei/go-dc/pkg/config"
	"log"
	"sync"
)

var mysqlOnce sync.Once
var mysqlMutex sync.Mutex
var mysqlClientMap map[string]*sql.DB

type IMySQL interface {
	DB1Client() *sql.DB
}

type MySQL struct {
	ClientMap map[string]*sql.DB
}

func BuildMySQL() IMySQL {
	initMySQL()
	return MySQL{
		ClientMap: mysqlClientMap,
	}
}

func (mysql MySQL) DB1Client() *sql.DB {
	return mysql.ClientMap["db1"]
}

func initMySQL() {
	mysqlConfigs := config.ServerConfig().MySQL
	mysqlOnce.Do(func() {
		mysqlMutex.Lock()
		defer mysqlMutex.Unlock()

		mysqlClientMap = make(map[string]*sql.DB, len(mysqlConfigs))
		for _, mysqlConfig := range mysqlConfigs {
			var err error
			var client *sql.DB
			if client, err = sql.Open(mysqlConfig.Type, mysqlConfig.Addr); err != nil {
				log.Fatalf("mysql open failed: %v", err)
				return
			}
			if err = client.Ping(); err != nil {
				log.Fatalf("mysql ping failed: %v", err)
				return
			}
			// 保存mysql客户端
			mysqlClientMap[mysqlConfig.Name] = client
		}
	})
}
