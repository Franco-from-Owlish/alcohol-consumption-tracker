package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	User     string
	Password string
	Port     string
	Host     string
	Name     string
	DBUrl    string
}

// NewDatabase
// Creates a new Database struct.
func NewDatabase(user, password, port, host, name string) *Database {
	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host, port, user, name, password,
	)
	return &Database{
		User:     user,
		Password: user,
		Port:     port,
		Host:     host,
		Name:     name,
		DBUrl:    dbURL,
	}
}

// Open Opens the connection to the database
func (db *Database) Open() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(db.DBUrl))
}
