package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func checkDNS(givenName string, outputName string) string {
	out, err := exec.Command("host", givenName).Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}

	output := string(out[:])

	if strings.Contains(output, outputName) {
		result := "Looks Good"
		fmt.Printf("|%-40s|%40s|%30s|\n", givenName, outputName, result)
	} else {
		result := "Oops Somethig Wrong"
		fmt.Printf("|%-40s|%40s|%30s|\n", givenName, outputName, result)
	}

	return output

}

func main() {
	checkDNS("google.com", "172.217.31.206")
	checkDNS("yahoo.com", "8.8.8.8")
	checkDNS("apple.com", "4.4.4.4")
	checkDNS("amazon.com", "1.1.1.1")
	checkDNS("microsoft.com", "104.215.148.63")
}
