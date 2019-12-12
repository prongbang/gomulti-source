package db

import (
	"fmt"
	"github.com/prongbang/gomulti-source/pkg/db/driver"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mariaDB struct {
	Cfg Config
}

func (db *mariaDB) Connect() *sqlx.DB {
	dsName := fmt.Sprintf("%s:%s@(%s:3306)/%s?charset=utf8&parseTime=true", db.Cfg.User, db.Cfg.Pass, db.Cfg.Host, db.Cfg.DatabaseName)
	conn, err := sqlx.Connect("mysql", dsName)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("MySQL Connected!")
	}
	return conn
}

// NewMariaDB for create mariadb driver
func NewMariaDB() driver.SQLxDriver {
	return &mariaDB{
		Cfg: Config{
			User:         "root",
			Pass:         "pass",
			Host:         "localhost",
			DatabaseName: "testDb",
		},
	}
}
