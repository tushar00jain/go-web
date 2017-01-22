package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/go-gorp/gorp"
	"html/template"
	"log"
	"net/http"
)

// schemas
type PhoneType int

const (
	MOBILE PhoneType = iota
	HOME
	WORK
)

type PhoneNumber struct {
	Number string
	// Type   PhoneType
	Type   int
}

type PhoneNumberArray struct {
	phoneNumbers []PhoneNumber
}

type Person struct {
	Id          int
	Name        string
	Email       string
	// PhoneNumberArray
	PhoneNumber
}

type AddressBook struct {
	People []Person
}

// template testing
type Page struct {
	Data string
}

func loadPage(data string) *Page {
	return &Page{Data: data}
}

// handler for index.html
func handler(w http.ResponseWriter, r *http.Request) {
	p := loadPage("go-web")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

// handler for getting persons
// func getPersons(db *sql.DB) func(http.ResponseWriter, *http.Request) {
func getPersons(dbmap *gorp.DbMap) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var data []Person
			_, err := dbmap.Select(&data, "SELECT * FROM Person")
			if err != nil {
				fmt.Println("selct error")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			result, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(result)
		} else {
			http.Error(w, "Wrong Method", http.StatusInternalServerError)
			fmt.Println("Wrong Method")
			return
		}
	}
}

func main() {
	db, err := sql.Open("postgres", "postgres://test:test@db/test?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// handle static files
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.FileServer(http.Dir("./static")))

	// api
	http.HandleFunc("/api/v1/persons", getPersons(dbmap))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
