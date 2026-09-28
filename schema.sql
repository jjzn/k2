PRAGMA foreign_keys = ON;

CREATE TABLE location (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT,
    lat REAL,
    lon REAL,
    metadata TEXT
) STRICT;

CREATE TABLE item (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    persons TEXT NOT NULL,
    location TEXT,
    description TEXT,
    date TEXT NOT NULL,
    enddate TEXT,
    endtime TEXT,
    isallday INT,
    FOREIGN KEY(location) REFERENCES location(id)
) STRICT;
