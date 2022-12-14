package db

import (
	"database/sql"
	"fmt"
	"log"

	"gihub.com/gadhittana01/product-project/config"
	"gihub.com/gadhittana01/product-project/helper"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func InitDB(dbname string) *sql.DB {
	config := &config.GlobalConfig{}
	helper.LoadConfig(config)

	sourceConn := config.DB[dbname]

	connString := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		sourceConn.Host, sourceConn.Port, sourceConn.User, sourceConn.Name,
	)

	db, err = sql.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB " + dbname + " Successfully connected!")

	return db
}
