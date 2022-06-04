# Item format docs

The `Item` structure holds all information relating to an individual item
(or entry) in the calendar:

```go
type Item struct {
	ID          string
        Title       string
        Persons     []string
        Description string
        Date        time.Time
        EndDate     time.Time
        EndTime     time.Time
        IsAllDay    bool
}
```

As `Item`s are stored in the `DB` database object, they implement a `key()`
function with the following signature:

```go
func (i Item) key() string
```

## A note about date representation

As the calendar supports both one-day items as well as items spanning multiple
days, dates in the `Item` structure use a quite idiosincratic format.

The `Date` member contains a `time.Time` representing the start time and date.
As the time component is not required for items, `Date.Hour()` and similar
functions may return a null value. In that case, the `IsAllDay` member is set
to `true` and consumers should only access the date component.

The end datetime is split into its two components (`EndDate` and `EndTime`).
Both are `time.Time` values, but only the date/time component is initialized,
so the same considerations as above apply for consumers.

If `EndDate` is uninitialized, `EndTime` (if defined) is assumed to happen on
the same date component as is described by `Date`. `EndTime` follows the same
logic, except for swapping the roles of the date and the time components.

### Why?

Because Go (and suboptimal design considerations). Mostly because `time.Time`
is too useful to be reimplemented containing only the date or time components,
and because the lack of algebraic data types (concretely, sum types and
optional types) make it hard to adapt this time type to the design requirements.

## Parsing items from requests

`Item`s can be obtained from a HTTP request by calling
`parseItem(w http.ResponseWriter, r *http.Request, id string) Item`. This
function takes the following arguments and constructs a new item from them:

- `w http.ResponseWriter`: the HTTP request's response handle, used for
reporting errors in the request.
- `r *http.Request`: the original HTTP request, for extracting item data.
- `id string`: the ID which will be assigned to the newly constructed item,
so that the caller has control over the ID generation scheme.
