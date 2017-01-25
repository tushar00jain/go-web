package utils

import (
	_ "github.com/lib/pq"
	"database/sql"
	"net/http"
)

func Only(method string, db *sql.DB, f func(db *sql.DB, w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			f(db, w, r)
			return
		}
		http.Error(w, method + " ONLY", http.StatusMethodNotAllowed)
		return
	}
}
