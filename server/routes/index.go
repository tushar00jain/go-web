package routes

import (
	"github.com/tushar00jain/go-web/server/schemas"
	"html/template"
	"net/http"
)

func loadPage(data string) *schemas.Page {
	return &schemas.Page{Data: data}
}

// handler for index.html
func Index(w http.ResponseWriter, r *http.Request) {
	p := loadPage("go-web")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}
