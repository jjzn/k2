/// Helper script to convert the app database from JSON (format as of Sept. 2026) to SQL

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// Taken from server.go
type Item struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Persons     []string  `json:"pers"`
	Location    string    `json:"loc"`
	LocationURL string    `json:"locurl"`
	Description string    `json:"desc"`
	Date        time.Time `json:"date"`
	EndDate     time.Time `json:"end_date"`
	EndTime     time.Time `json:"end_time"`
	IsAllDay    bool      `json:"all_day"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s input.json output.sql\n", os.Args[0])
		return
	}

	infile := os.Args[1]
	outfile := os.Args[2]

	raw, err := os.ReadFile(infile)
	if err != nil {
		panic(err)
	}

	var jsondb map[string]Item
	if err := json.Unmarshal(raw, &jsondb); err != nil {
		panic(err)
	}

	db, err := sql.Open("sqlite3", outfile)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// TODO: check DB is empty

	ins_item, err := db.Prepare("INSERT INTO item(id, title, persons, location, description, date, enddate, endtime, isallday) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer ins_item.Close()

	ins_loc, err := db.Prepare("INSERT INTO location(id, name, url) VALUES(?, ?, ?)");
	if err != nil {
		panic(err)
	}
	defer ins_loc.Close()

	for _, item := range jsondb {
		if item.Location != "" {
			// TODO: avoid duplicate locations (sharing same name)
			_, err := ins_loc.Exec(uuid.New().String(), item.Location, item.LocationURL)
			if err != nil {
				panic(err)
			}
		}

		// TODO: how should []string be represented? for item.Persons
		_, err := ins_item.Exec(item.ID, item.Title, strings.Join(item.Persons, ", "), item.Location, item.Description, item.Date, item.EndDate, item.EndTime, item.IsAllDay)
		if err != nil {
			panic(err)
		}
	}
}
