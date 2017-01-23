package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tushar00jain/go-web/server/schemas"
	"net/http"
)

// handler for getting persons
func GetPersons(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			rows, err := db.Query("SELECT p.\"Id\", p.\"Name\", p.\"Email\"," +
				"json_agg(json_build_object('Number', ph.\"Number\", 'Type', ph.\"Type\")) AS \"PhoneNumbers\"" +
				"FROM Person p INNER JOIN PhoneNumber ph ON p.\"Id\"=ph.\"PersonId\" GROUP BY p.\"Id\";")

			if err != nil {
				fmt.Println("selct error")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := make([]schemas.Person, 0)
			for rows.Next() {
        var person schemas.Person
				var temp string
				err = rows.Scan(&person.Id, &person.Name, &person.Email, &temp)

				if err != nil {
					fmt.Println("query error")
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				tbs := []byte(temp)
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

func GetAddressBook(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			var query = "SELECT d.*, json_agg(json_build_object('Number', ph.\"Number\", 'Type', ph.\"Type\")) AS \"PhoneNumbers\" FROM (" +
										"SELECT p.*" +
										"FROM Person p INNER JOIN AddressBook a ON p.\"Id\"=a.\"People\" AND a.\"Self\"=" + r.URL.Query().Get("id") +
										") d INNER JOIN PhoneNumber ph ON d.\"Id\"=ph.\"PersonId\" GROUP BY d.\"Id\", d.\"Name\", d.\"Email\";"
			rows, err := db.Query(query)
			if err != nil {
				fmt.Println("selct error")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := make([]schemas.Person, 0)
			for rows.Next() {
        var person schemas.Person
				var temp string
				err = rows.Scan(&person.Id, &person.Name, &person.Email, &temp)

				if err != nil {
					fmt.Println("query error")
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				tbs := []byte(temp)
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
