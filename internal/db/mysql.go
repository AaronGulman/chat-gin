package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARN] Could not load .env file, using system env")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("[FATAL] DB_URL not set in environment")
	}

	var err error
	DB, err = sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatalf("[FATAL] Failed to open database connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("[FATAL] Could not connect to the database: %v", err)
	}

	DB.SetConnMaxLifetime(3 * time.Minute)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	log.Println("[INFO] Database connection successfully initialized")
}
