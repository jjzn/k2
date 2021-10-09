package k2

import (
	"sort"
	"strconv"
	"time"

	"encoding/json"
	rt "github.com/julienschmidt/httprouter"
	"net/http"
)

func serveAPI(filter func(Item) bool, w http.ResponseWriter) {
	items := data.Filter(filter)
	sort.Slice(items, sortByDate(items))

	raw, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(raw)
	w.Write([]byte("\n"))
}

func apiAll(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	serveAPI(func(i Item) bool {
		return today.Before(i.Date)
	}, w)
}

func apiDay(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	year, err := strconv.Atoi(ps.ByName("year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	month, err := strconv.Atoi(ps.ByName("month"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	day, err := strconv.Atoi(ps.ByName("day"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serveAPI(func(i Item) bool {
		return i.Date.Year() == year &&
			int(i.Date.Month()) == month &&
			i.Date.Day() == day
	}, w)
}

func apiMonth(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	year, err := strconv.Atoi(ps.ByName("year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	month, err := strconv.Atoi(ps.ByName("month"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	serveAPI(func(i Item) bool {
		return i.Date.Year() == year &&
			int(i.Date.Month()) == month
	}, w)
}

func apiItem(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	item, ok := data.Get(ps.ByName("id"))
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	raw, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(raw)
	w.Write([]byte("\n"))
}
