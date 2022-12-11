package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	rt "github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var fns = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
	"formatTime": formatTime,
	"join": strings.Join,
	"formatDates": formatDates,
}

var list = template.Must(
	template.New("list").Funcs(fns).ParseFiles("templ/layout", "templ/list"))
var grid = template.Must(
	template.New("grid").Funcs(fns).ParseFiles("templ/layout", "templ/grid"))
var details = template.Must(
	template.New("details").Funcs(fns).ParseFiles("templ/layout", "templ/details"))
var invite = template.Must(
	template.New("invite").Funcs(fns).ParseFiles("templ/layout", "templ/invite"))

func sortByDate(items []Item) func(int, int) bool {
	return func(i, j int) bool {
		if items[i].Date.Equal(items[j].Date) {
			return items[i].Title < items[j].Title
		} else {
			return items[i].Date.Before(items[j].Date)
		}
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	items := data.Filter(func(i Item) bool {
		return isInDateRange(i, today) || today.Before(i.Date)
	})

	sort.Slice(items, sortByDate(items))

	if err := list.Execute(w, items); err != nil {
		panic(err)
	}
}

func handleDate(now time.Time, w http.ResponseWriter) {
	items := data.Filter(func(i Item) bool {
		return isInDateRange(i, now)
	})

	sort.Slice(items, sortByDate(items))

	if err := list.Execute(w, items); err != nil {
		panic(err)
	}
}

func handleToday(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	handleDate(time.Now(), w)
}

func handleTomorrow(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	handleDate(time.Now().AddDate(0, 0, 1), w)
}

func handleWeek(year, week int, w http.ResponseWriter) {
	start := firstDayOfWeek(year, week)

	items := data.Filter(func(i Item) bool {
		for k := 0; k < 7; k++ {
			day := start.AddDate(0, 0, k)

			if isInDateRange(i, day) {
				return true
			}
		}

		return false
	})

	sort.Slice(items, sortByDate(items))

	if err := list.Execute(w, items); err != nil {
		panic(err)
	}
}

func handleThisWeek(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	year, week := time.Now().ISOWeek()
	handleWeek(year, week, w)
}

func handleNextWeek(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	nw := time.Now().AddDate(0, 0, 7)
	year, week := nw.ISOWeek()
	handleWeek(year, week, w)
}

func handleMonth(year int, month time.Month, w http.ResponseWriter) {
	now := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().Location())
	n_days := now.AddDate(0, 1, -1).Day()

	days := make([][]Item, n_days)
	items := data.Filter(func(i Item) bool {
		for k := 0; k < n_days; k++ {
			if isInDateRange(i, now.AddDate(0, 0, k)) {
				return true
			}
		}

		return false
	})

	sort.Slice(items, sortByDate(items))

	for _, i := range items {
		days[i.Date.Day()-1] = append(days[i.Date.Day()-1], i)
	}

	d := struct {
		Month  string
		Filler []int
		Days   [][]Item
	}{
		formatMonth(month),
		make([]int, (now.Weekday()+6)%7),
		days,
	}

	if err := grid.Execute(w, d); err != nil {
		panic(err)
	}
}

func handleFilterMonth(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	year, err := strconv.Atoi(ps.ByName("year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "cannot parse year\n")
		return
	}

	date, err := time.Parse("1", ps.ByName("month"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "cannot parse month\n")
		return
	}

	handleMonth(year, date.Month(), w)
}

func handleThisMonth(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	year, month, _ := time.Now().Date()
	handleMonth(year, month, w)
}

func handleNextMonth(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	nw := time.Now().AddDate(0, 1, 0)
	year, month, _ := nw.Date()
	handleMonth(year, month, w)
}

func handleView(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	item, ok := data.Get(ps.ByName("id"))
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "item not found\n")
		return
	}

	if err := details.Execute(w, item); err != nil {
		panic(err)
	}
}

func handleInvite(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	item, ok := data.Get(ps.ByName("id"))
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invite not found\n")
		return
	}

	if err := invite.Execute(w, item); err != nil {
		panic(err)
	}
}

func handleFilterDate(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	year, err := strconv.Atoi(ps.ByName("year"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "cannot parse year\n")
		return
	}

	day, err := strconv.Atoi(ps.ByName("day"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "cannot parse day\n")
		return
	}

	date, err := time.Parse("1", ps.ByName("month"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "cannot parse month\n")
		return
	}

	now := time.Date(year, date.Month(), day, 0, 0, 0, 0, time.Now().Location())

	items := data.Filter(func(i Item) bool {
		return isInDateRange(i, now)
	})

	sort.Slice(items, sortByDate(items))

	if err := list.Execute(w, items); err != nil {
		panic(err)
	}
}
