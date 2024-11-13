package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type Document map[string]interface{}

type Database struct {
	data map[string]Document
	mu   sync.RWMutex
	file string
}

// Loads data from a file in memory or starts fresh
func (db *Database) load() {
	file, err := os.OpenFile(db.file, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("No existing file found, starting fresh...")
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()

	var data map[string]Document
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		db.data = data
	}
}

// Init - Initialize database
func Init(file string) *Database {
	db := &Database{
		data: make(map[string]Document),
		file: file,
	}
	db.load()

	return db
}
