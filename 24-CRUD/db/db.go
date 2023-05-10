package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	strConnection := "root:ilewml@/golangCRUD?charset=utf8&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
