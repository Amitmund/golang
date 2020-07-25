package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

/*

Quick Notes:

// Defining the main function.
// The used functions are defined below of the main function.


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
