package models

import (
	"os"
	"path"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_NAME = "data/voyager.db"
	_SQL_DRIVER = "sqlite3"
)

func InitDB() {
	// create DB file and path folder
	if _, err := os.Stat(_DB_NAME); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)

		if err != nil {
			panic(err)
		}
	}

	orm.RegisterModel(new(Post), new(Comment), new(User))

	orm.RegisterDriver(_SQL_DRIVER, orm.DRSqlite)

	orm.RegisterDataBase("default", _SQL_DRIVER, _DB_NAME, 10)
}
