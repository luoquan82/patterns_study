package singleton

import (
	"fmt"
	"sync"
)

type DB struct{}

var db *DB
var once sync.Once

func InitDB(dsn string) *DB {
	fmt.Println("init db with dsn:", dsn)
	return &DB{}
}

func GetInstance() *DB {
	once.Do(func() {
		db = InitDB("")
	})
	return db
}
