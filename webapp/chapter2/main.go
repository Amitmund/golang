package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

//
type NewsAggPage struct {
	Title string
	News  string
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "some news"}
	t, _ := template.ParseFiles("basictemplating.html")
	t.Execute(w, p)
}

//
type DataCenter struct {
	DC string
}

func dnsCheck(w http.ResponseWriter, r *http.Request) {
	p := DataCenter{DC: "p3"}
	t, _ := template.ParseFiles("dnsCheck.html")
	t.Execute(w, p)
	// nil also work
	// t.Execute(w, nil)
	return

	details := DataCenter {
		DC := t.FormValue("dc")
	}

	// do something.
	fmt.Println("dc")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)
	http.HandleFunc("/dnsCheck/", dnsCheck)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	http.ListenAndServe(":8000", nil)
}

/*
// forms.go
package main


type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

func main() {
    tmpl := template.Must(template.ParseFiles("forms.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // do something with details
        _ = details

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}
*/
