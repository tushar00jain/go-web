package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tushar00jain/go-web/server/routes"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "postgres://test:test@db/test?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// handle static files
	http.HandleFunc("/", routes.Index)
	http.Handle("/static/", http.FileServer(http.Dir("./static")))

	// api
	http.HandleFunc("/api/v1/persons", routes.GetPersons(db))
	http.HandleFunc("/api/v1/addressBook", routes.GetAddressBook(db))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
