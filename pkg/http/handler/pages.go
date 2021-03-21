package handler

import (
	"html/template"
	"net/http"

	"github.com/daBONDi/go-rest-wol/internal/repository"
)

// RenderHomePage The index page
func RenderHomePage(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("pages/index.html")
	tmpl.Execute(w, repository.ComputerList)

}
