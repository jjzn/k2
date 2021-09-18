# API docs

## Response Format

The response format for all API requests is JSON. It can either
be an array of items, or a single item.

The format for an item is:

```json
{
    "id": "[ID]",
    "title": "[TITLE]",
    "author": "[AUTHOR]",
    "desc": "[DESCRIPTION]",
    "date": "[ISO 8601 DATE]",
    "end": "[ISO 8601 DATE, MAY BE NULL]",
    "all_day": true/false
}
```

## Endpoints

Base URL for the v1 API: `api.k2.test/v1`

### `/all`

`/all` returns a list of all calendar entries in JSON format.

### `/date/:year/:month`

`/date/:year/:month` returns a list of all calendar entries for
the specified year and month in JSON format.

- `year` is a 4-digit integer between 0 and 9999.
- `month` is a 1-digit or 2-digit integer between 1 and 12.

### `/date/:year/:month/:day`

`/date/:year/:month/:day` returns a list of all calendar entries
for the specified year, month and day in JSON format.

- `year` is a 4-digit integer between 0 and 9999.
- `month` is a 1-digit or 2-digit integer between 1 and 12.
- `day` is a 1-digit or 2-digit integer between 1 and 28/29/30/31
(depending on the month).

### `/id/:id`

`id/:id` returns the calendar entry with the specified ID.
`id` is a UUIDv4 in the `XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX`
string format.
