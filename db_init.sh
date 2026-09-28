#!/bin/sh

dbname="${1:-app.sqlite.db}"
echo "$dbname: Creating DB with schema"
sqlite3 "$dbname" < "$(dirname $0)/schema.sql"

#echo "$dbname: Populating DB with mock data"
#sqlite3 "$dbname" < "$(dirname $0)/populate_mock.sql"
