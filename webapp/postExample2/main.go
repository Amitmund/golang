package main

import (
	"dnscheck"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// main functions:

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/dnsCheck", dnsCheck)
	http.HandleFunc("/subject", subject)
	http.HandleFunc("/commands", commands)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// hello function

func hello(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.Error(w, "404 not found.", http.StatusNotFound)
	// 	return
	// }

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

// for dnscheck

func dnsCheck(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "dnsCheck.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from dnsCheck functions! r.PortForm = %v\n\n", r.PostForm)

		data := dnscheck.DNSCheck("portal-db.akamai.com", "prod-portal-db.akamai.com")
		fmt.Fprintf(w, "portal-db.akamai.com ==> ", "prod-portal-db.akamai.com ==> ", data)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

// for subject

func subject(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "subject.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from subject functions! r.PortForm = %v\n", r.PostForm)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

// for commands

func commands(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "commands.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from subject functions! r.PortForm = %v\n", r.PostForm)
		command := r.FormValue("command")
		cmdOutput, _ := exec.Command("/bin/bash", "-c", command).Output()
		stringOutput := string(cmdOutput[:])
		fmt.Fprintf(w, stringOutput)

		// 	stringOutput := string(cmdOutput[:])
		// fmt.Fprintf(w, exec.Command())
		// name := r.FormValue("name")
		// address := r.FormValue("address")
		// fmt.Fprintf(w, "Name = %s\n", name)
		// fmt.Fprintf(w, "Address = %s\n", address)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}
