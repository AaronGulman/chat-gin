package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB
var err error

func Init() {
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		fmt.Println("Error loading .env file", err)
	}
	fmt.Println("This is the url: ", DB_URL)

	DB, err = sql.Open("mysql", DB_URL)
	if err != nil {
		panic(err)
	}

	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

}
