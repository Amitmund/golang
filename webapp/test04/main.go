package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// To explore on net/http/Request
	// https://golang.org/pkg/net/http/#Request

	fmt.Fprintf(w, "To understand the http.Request internals\n")

	// Just to give a new line in the output
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  Method ==> %s", r.Method)

	//

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "\n")

	// r.URL.Path to provide url path.
	fmt.Fprintf(w, "To see ==>  URL.Path[1:] ==> %s", r.URL.Path[1:])

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  URL.Path ==> %s", r.URL.Path)

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  URL.User, for Userinfo ==> %s", r.URL.User)

	//

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  RequestURI ==> %s", r.RequestURI)

	//

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  Proto ==> %s", r.Proto)

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  ProtoMajor (int data type) ==> %d", r.ProtoMajor)

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  ProtoMinor (int data type) ==> %d", r.ProtoMinor)

	//

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  Header (a map / Key:Value pair(s)) ==> %s", r.Header)

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  Header (a Key:Value of User-Agent) ==> %s", r.Header["User-Agent"])

	//

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  Host ==> %s", r.Host)

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  RemoteAddr ==> %s", r.RemoteAddr)

	//

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "To see ==>  ContentLength ==> %d", r.ContentLength)

	//

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)

	fmt.Println("Web Server")
}
