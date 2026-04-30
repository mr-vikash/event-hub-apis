package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "root:root@tcp(127.0.0.1:3311)/eventhub"

	var err error

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("DB open error:", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("DB ping error", err)
	}

	log.Println("Connected to DB")
}
