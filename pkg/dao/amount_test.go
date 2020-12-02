package dao

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/zhulinwei/go-dc/pkg/model"
)

func NewDBMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestAmountDao_SaveAmount(t *testing.T) {
	db, mock := NewDBMock()
	amountDao := AmountDao{dbClient: db}
	mockUserAmount := model.UserAmountRequest{Name: "tony", Amount: 1000}

	query := "insert into user_amount \\(name, amount\\) values \\(\\?, \\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(mockUserAmount.Name, mockUserAmount.Amount).WillReturnResult(sqlmock.NewResult(0, 0))

	_, err := amountDao.SaveAmount(mockUserAmount)
	assert.NoError(t, err)
}
