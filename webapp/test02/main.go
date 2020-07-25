package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func staticPage1(w http.ResponseWriter, r *http.Request) {
	// file name should be with in ""
	// Note: the template, can't display images or so...
	staticPage1, _ := template.ParseFiles("pages/staticPage1.html")
	staticPage1.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/staticPage1", staticPage1)
	http.ListenAndServe(":80", nil)
}
