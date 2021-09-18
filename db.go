package main

import (
	"encoding/json"
	"io/ioutil"
)

type DB struct {
	File string
	Map  map[string]Item
}

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

func (db *DB) Set(key string, item Item) {
	if db.Map == nil {
		db.Map = make(map[string]Item)
	}

	db.Map[key] = item
}

func (db *DB) Get(key string) (Item, bool) {
	item, ok := db.Map[key]
	return item, ok
}

func (db *DB) Insert(item Item) string {
	if db.Map == nil {
		db.Map = make(map[string]Item)
	}

	db.Map[item.key()] = item
	return item.key()
}

func (db *DB) Delete(key string) {
	delete(db.Map, key)
}

func (db *DB) Filter(fn func(Item) bool) []Item {
	items := make([]Item, 0, len(db.Map))
	for _, item := range db.Map {
		if fn(item) {
			items = append(items, item)
		}
	}

	return items
}

func (db *DB) ForEach(fn func(Item)) {
	for _, item := range db.Map {
		fn(item)
	}
}

func (db *DB) Items() []Item {
	return db.Filter(func(_ Item) bool { return true })
}
