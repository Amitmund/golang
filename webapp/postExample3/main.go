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
		// Commenting the following line, not to display the post details.
		// fmt.Fprintf(w, "Post from subject functions! r.PortForm = %v\n", r.PostForm)
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
	http.HandleFunc("/", auth.JustCheck(authenticator, commands))

	// you this authenticator, only if needed. else you can just use the simple one.
	// Example:
	// http.HandleFunc("/", hello)
	// http.HandleFunc("/commands", auth.JustCheck(authenticator, commands))
	// http.HandleFunc("/commands/", auth.JustCheck(authenticator, commands))

	fmt.Printf("Welcome to Amit Mund's  commandsOverUI...\n")
	fmt.Printf("Starting server at port 80...\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

// htpasswd -c .htpasswd amitmund
// To add a new user "htpasswd .htpasswd devops"
// note: the .htpasswd file should be there on the same location, where the code is running.
// 	auth "github.com/abbot/go-http-auth"
// htpasswd := auth.HtpasswdFileProvider("./.htpasswd")
// authenticator := auth.NewBasicAuthenticator("Basic Realm", htpasswd)
// http.HandleFunc("/commands/", auth.JustCheck(authenticator, commands))
