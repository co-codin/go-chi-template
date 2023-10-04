package db

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute


func ConnectPostgres(dsn string) (*DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifeTime)

	err = testDB(db)

	if err != nil {
		return nil, err
	}

	dbConn.DB = db

	return dbConn, nil
}

func testDB(d *sql.DB) error {
	err := d.Ping()

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("Success: Database is connected")
	return nil
}