package app

import (
	"database/sql"
	"kasirku/helpers"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DB() *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("CONS_DB")
	if dsn == "" {
		log.Fatal("DB Connetion Is Not Exist")
	}

	db, err := sql.Open("postgres", dsn)
	helpers.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
