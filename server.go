package main

import (
	"html/template"
	"net/http"

	"flag"
	"sort"
	"strings"
	"time"

	"fmt"
	"log"

	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	rt "github.com/julienschmidt/httprouter"
)

type Item struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Persons     []string  `json:"pers"`
	Description string    `json:"desc"`
	Date        time.Time `json:"date"`
	EndDate     time.Time `json:"end_date"`
	EndTime     time.Time `json:"end_time"`
	IsAllDay    bool      `json:"all_day"`
}

var data DB

var entryForm = template.Must(
	template.New("new").Funcs(fns).ParseFiles("templ/layout", "templ/new"))

func (i Item) key() string {
	return i.ID
}

func parseItem(w http.ResponseWriter, r *http.Request, id string) Item {
	failBadRequest := func(msg string) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, msg)
	}

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	title := strings.TrimSpace(r.PostForm.Get("title"))
	persons := strings.Split(r.PostForm.Get("persons"), ",")

	for i, elem := range persons {
		persons[i] = strings.TrimSpace(elem)
	}

	sort.Strings(persons)

	if title == "" || len(persons) == 0 {
		failBadRequest("missing field (title, persons)\n")
		return Item{}
	}

	desc := r.PostForm.Get("desc")

	all_day := true

	f := "2006-01-02 "
	if r.PostForm.Get("time") != "" {
		f += "15:04"
		all_day = false
	}

	dt := r.PostForm.Get("date") + ` ` + r.PostForm.Get("time")
	date, err := time.Parse(f, dt)
	if err != nil {
		failBadRequest("cannot parse date\n")
		return Item{}
	}

	end_date, err := time.Parse("2006-01-02", r.PostForm.Get("end-date"))
	if err != nil {
		end_date = time.Time{}
	}

	end_time, err := time.Parse("15:04", r.PostForm.Get("end-time"))
	if err != nil {
		end_time = time.Time{}
	}

	return Item{id, title, persons, desc, date, end_date, end_time, all_day}
}

func handleAdd(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	id := uuid.New().String()
	item := parseItem(w, r, id)
	if item.Title == "" {
		return
	}

	data.Insert(item)
	http.Redirect(w, r, "/view/"+id, http.StatusSeeOther)
}

func handleUpdate(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	id := ps.ByName("id")
	item := parseItem(w, r, id)
	if item.Title == "" {
		return
	}

	data.Set(id, item)
	http.Redirect(w, r, "/view/"+id, http.StatusSeeOther)
}

func handleAddPerson(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	id := ps.ByName("id")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	name := r.PostForm.Get("name")
	item, ok := data.Get(id)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "item not found\n")
		return
	}

	if name != "" {
		item.Persons = append(item.Persons, name)
		data.Set(id, item)
	}

	http.Redirect(w, r, "/view/"+id, http.StatusSeeOther)
}

func handleNew(w http.ResponseWriter, r *http.Request, _ rt.Params) {
	if err := entryForm.Execute(w, Item{}); err != nil {
		panic(err)
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	data.Delete(ps.ByName("id"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleEdit(w http.ResponseWriter, r *http.Request, ps rt.Params) {
	item, ok := data.Get(ps.ByName("id"))
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "item not found\n")
		return
	}

	if err := entryForm.Execute(w, item); err != nil {
		panic(err)
	}
}

func main() {
	log.Println("opening database")
	d, err := OpenDB("data.json")
	if err != nil {
		panic(err)
	}
	data = d

	t := time.NewTicker(12 * time.Hour)
	go func() {
		for _ = range t.C {
			log.Println("writing periodic backup")
			if err := data.Write(); err != nil {
				panic(err)
			}
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("terminating... writing data")
		if err := data.Write(); err != nil {
			panic(err)
		}
		os.Exit(0)
	}()

	r := rt.New()

	r.POST("/add", handleAdd)
	r.POST("/update/:id", handleUpdate)
	r.POST("/add-person/:id", handleAddPerson)

	r.GET("/new", handleNew)
	r.GET("/delete/:id", handleDelete)
	r.GET("/view/:id", handleView)
	r.GET("/invite/:id", handleInvite)
	r.GET("/edit/:id", handleEdit)

	r.GET("/date/:year/:month/:day", handleFilterDate)
	r.GET("/date/:year/:month", handleFilterMonth)
	r.GET("/week/:year/:week", handleFilterWeek)

	r.GET("/today", handleToday)
	r.GET("/tomorrow", handleTomorrow)
	r.GET("/this-week", handleThisWeek)
	r.GET("/next-week", handleNextWeek)
	r.GET("/this-month", handleThisMonth)
	r.GET("/next-month", handleNextMonth)

	r.GET("/api/v1/all", apiAll)
	r.GET("/api/v1/date/:year/:month", apiMonth)
	r.GET("/api/v1/date/:year/:month/:day", apiDay)
	r.GET("/api/v1/id/:id", apiItem)

	r.GET("/", handleIndex)

	r.ServeFiles("/static/*filepath", http.Dir("static"))
	r.ServeFiles("/docs/*filepath", http.Dir("docs"))

	port := flag.String("p", "8020", "port number")
	flag.Parse()

	log.Println("listening on port", *port)
	log.Fatal(http.ListenAndServe(":"+*port, r))
}
