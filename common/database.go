package common

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database struct {
	*sql.DB
}

// var DB *gorm.DB
var DB *sql.DB

// Opening a database and save the reference to `Database` struct.
func Init() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)

	DB = db
	return DB
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *sql.DB {
	return DB
}
