package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	// "reflect"
	_ "github.com/lib/pq"
	// "github.com/go-gorp/gorp"
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
	Number string `json:"Number"`
	Type   PhoneType `json:"Type"`
}

type Person struct {
	Id          int
	Name        string
	Email       string
	PhoneNumbers []PhoneNumber
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
func getPersons(db *sql.DB) func(http.ResponseWriter, *http.Request) {
// func getPersons(dbmap *gorp.DbMap) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// var data []Person
			rows, err := db.Query("SELECT p.\"Id\", p.\"Name\", p.\"Email\"," +
				"json_agg(json_build_object('Number', ph.\"Number\", 'Type', ph.\"Type\")) AS \"PhoneNumbers\"" +
				"FROM Person p INNER JOIN PhoneNumber ph ON p.\"Id\"=ph.\"PersonId\" GROUP BY p.\"Id\";")
			if err != nil {
				fmt.Println("selct error")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			data := make([]Person, 0)
			for rows.Next() {
        var person Person
				var temp string
				err = rows.Scan(&person.Id, &person.Name, &person.Email, &temp)
				if err != nil {
					fmt.Println("query error")
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				tbs := []byte(temp)
				// phoneNumbers := make([]PhoneNumber, 0)
				if err := json.Unmarshal(tbs, &person.PhoneNumbers); err != nil {
					fmt.Println("unmarshal error")
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				data = append(data, person)
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

	// dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	// handle static files
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.FileServer(http.Dir("./static")))

	// api
	http.HandleFunc("/api/v1/persons", getPersons(db))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
