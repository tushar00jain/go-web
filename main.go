package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
)

// schemas
type PhoneType int

const (
	MOBILE PhoneType = iota
	HOME
	WORK
)

type PhoneNumber struct {
	phoneNumber string
	phoneType PhoneType
}

type Person struct {
	id int32 `json:"id,omitempty"`
	name string `json:"name,omitempty"`
	email string `json:"email,omitempty"`
	phoneNumber PhoneNumber `json:"phoneNumber,omitempty"`
}

type AddressBook struct {
	self Person `json:"self,omitempty"`
	people Person `json:"people,omitempty"`
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
func getPersons(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			data := Person{1, "test", "test@test.com", PhoneNumber{"5555555555", MOBILE}}
			result, err := json.Marshal(data)
			if err != nil {
				fmt.Println("JSON")
			}
			fmt.Println("DB")
			w.Header().Set("Content-Type", "application/json")
			w.Write(result)
			fmt.Println(result)
		}
			fmt.Println("Error")
	}
}

func main() {
	db, err := sql.Open("postgres", "test:test@db/test")
	if err != nil {
    log.Fatal(err)
	}
	defer db.Close()

	// handle static files
  http.HandleFunc("/", handler)
	http.Handle("/static/", http.FileServer(http.Dir("./static")))

	// api
  http.HandleFunc("/api/v1/persons", getPersons(db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
