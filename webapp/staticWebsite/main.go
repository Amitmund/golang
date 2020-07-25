package main

/*

Details from the following url:
https://tutorialedge.net/golang/creating-simple-web-server-with-golang/

Interesting one is:

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

As long as the static file is there, and you type the same static file name at URI, it will render and serve that page.

example:
localhost:8081/test.html ( then make sure, you have a page called test.html)
localhost:8081/hi.html ( then make sure, you have a page called hi.html)

*/

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
