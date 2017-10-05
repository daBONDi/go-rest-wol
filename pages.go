package main

import (
	"html/template"
	"net/http"
)

func renderHomePage(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("pages/index.html")
	tmpl.Execute(w, ComputerList)

}
