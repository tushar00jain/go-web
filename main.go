package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/tushar00jain/go-web/server/routes"
	"github.com/tushar00jain/go-web/server/utils"
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
	http.HandleFunc("/api/v1/persons", utils.Only("GET", db, routes.GetPersons))
	http.HandleFunc("/api/v1/addressBook", utils.Only("GET", db, routes.GetAddressBook))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
