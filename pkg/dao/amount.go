package dao

import (
	"database/sql"
	"fmt"

	"github.com/zhulinwei/go-dc/pkg/database"
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util"
	"github.com/zhulinwei/go-dc/pkg/util/log"
)

type IAmountDao interface {
	SaveAmount(userAmount model.UserAmountRequest) (*sql.Result, error)
}

type AmountDao struct {
	dbClient *sql.DB
}

func BuildAmountDao() IAmountDao {
	fmt.Println(database.BuildMySQL().DBClient())
	return AmountDao{
		dbClient: database.BuildMySQL().DBClient(),
	}
}

func (amountDao AmountDao) SaveAmount(userAmount model.UserAmountRequest) (*sql.Result, error) {
	query := "insert into user_amount (name, amount) values (?, ?)"
	stmt, err := amountDao.dbClient.PrepareContext(util.CommonContent(), query)
	if err != nil {
		log.Error("amount dao save amount prepare fail", log.String("error", err.Error()))
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(util.CommonContent(), userAmount.Name, userAmount.Amount)
	if err != nil {
		log.Error("amount dao save amount exec fail", log.String("error", err.Error()))
		return nil, err
	}

	return &result, nil
}
