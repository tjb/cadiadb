package internal

import (
	"encoding/json"
	"fmt"
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
	file, err := os.Open(db.file)

	if err != nil {
		fmt.Println("No existing file found, starting fresh...")
	}

	defer file.Close()

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
