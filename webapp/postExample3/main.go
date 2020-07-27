package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	auth "github.com/abbot/go-http-auth"
)

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

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

// main functions:

func main() {

	// read from .htpasswd file
	htpasswd := auth.HtpasswdFileProvider("./.htpasswd")
	authenticator := auth.NewBasicAuthenticator("Basic Realm", htpasswd)

	// http.HandleFunc("/commands", commands)
	http.HandleFunc("/commands", auth.JustCheck(authenticator, commands))
	http.HandleFunc("/commands/", auth.JustCheck(authenticator, commands))

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// htpasswd -c .htpasswd amitmund
// note: the .htpasswd file should be there on the same location, where the code is running.
// 	auth "github.com/abbot/go-http-auth"
// htpasswd := auth.HtpasswdFileProvider("./.htpasswd")
// authenticator := auth.NewBasicAuthenticator("Basic Realm", htpasswd)
// http.HandleFunc("/commands/", auth.JustCheck(authenticator, commands))
