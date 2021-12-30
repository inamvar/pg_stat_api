package database

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var Db *sql.DB




func CreateMockConnection() sqlmock.Sqlmock{
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	Db = db
	return  mock
}

// create connection with postgres db
func CreateConnection() (*sql.DB, error) {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		return nil, err
	}

	if db == nil {
		return nil, errors.New("database connection failed")
	}
	// check the connection
	err = db.Ping()

	if err != nil {
		return nil, err
	}
	Db = db
	// return the connection
	return db, nil

}
