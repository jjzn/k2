package main

import (
	"encoding/json"
	"io/ioutil"
)

// DB represents a persisted in-memory key-value database.
type DB struct {
	File string
	Map  map[string]Item
}

// OpenDB tries to open a named file as a DB database.
// The data in the file must be in JSON format.
func OpenDB(file string) (DB, error) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return DB{}, err
	}

	var db map[string]Item
	if err := json.Unmarshal(raw, &db); err != nil {
		return DB{}, err
	}

	if db == nil {
		db = make(map[string]Item)
	}

	return DB{file, db}, nil
}

// Write flushes the database to the associated file on disk.
func (db *DB) Write() error {
	raw, err := json.Marshal(db.Map)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(db.File, raw, 0644); err != nil {
		return err
	}

	return nil
}

// Set changes the item associated with the given key in the
// database. If no such entry exists, it will be created.
func (db *DB) Set(key string, item Item) {
	if db.Map == nil {
		db.Map = make(map[string]Item)
	}

	db.Map[key] = item
}

// Get returns the item associated with the given key.
func (db *DB) Get(key string) (Item, bool) {
	item, ok := db.Map[key]
	return item, ok
}

// WARNING: DEPRECATED FUNCTION
// Insert takes an item, computes its key and stores it in the
// database. It returns the key associated the item.
func (db *DB) Insert(item Item) string {
	if db.Map == nil {
		db.Map = make(map[string]Item)
	}

	db.Map[item.key()] = item
	return item.key()
}

// Delete removes the database entry with the given key.
func (db *DB) Delete(key string) {
	delete(db.Map, key)
}

// Filter iterates over all database items and returns those for
// which the filter function returns true.
func (db *DB) Filter(fn func(Item) bool) []Item {
	items := make([]Item, 0, len(db.Map))
	for _, item := range db.Map {
		if fn(item) {
			items = append(items, item)
		}
	}

	return items
}

// ForEach loops over every items in the database and executes
// a function on it. The function should not add new items to the
// database, as it cannot be guaranteed that the function will be
// run on it. The iteration order is unstable and should not be
// relied on.
func (db *DB) ForEach(fn func(Item)) {
	for _, item := range db.Map {
		fn(item)
	}
}

// Items returns all items currently in the database. There is
// no guaranteed sort order.
func (db *DB) Items() []Item {
	return db.Filter(func(_ Item) bool { return true })
}
