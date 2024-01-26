package db

import (
    "fmt"
    "log"
    "os"
    "github.com/jmoiron/sqlx"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

var db *sqlx.DB

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")

    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbName)
    db, err = sqlx.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatal(err)
    }

    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to the database")
}

func GetDB() *sqlx.DB {
    return db
}