package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type DataStore struct {
	User     string
	Password string
	Port     string
	Host     string
	Name     string
	Address  string
}

// NewDataStore
// Creates a new DataStore struct.
func NewDataStore(user, password, port, host, name string) *DataStore {
	dbURL := fmt.Sprintf("%s:%s", host, port)
	return &DataStore{
		User:     user,
		Password: password,
		Port:     port,
		Host:     host,
		Name:     name,
		Address:  dbURL,
	}
}

// Open Opens the connection to the database
func (db *DataStore) Open() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:        db.Address,
		Username:    db.User,
		Password:    db.Password,
		DB:          0,
		ReadTimeout: -1,
	})
	status, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("redis connection was refused: %v \nds: %v\n", err, db))
	}
	fmt.Println(status)
	return client
}
