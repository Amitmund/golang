package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

// Dns check.
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
		fmt.Printf("|%-40s|%30s|%30s|\n", givenName, outputName, result)
	} else {
		result := "Oops Somethig Wrong"
		fmt.Printf("|%-40s|%30s|%30s|\n", givenName, outputName, result)
	}

	return output

}

// to put the underline at the output.
func underline() {
	fmt.Printf("-------------------------------------------------------------------------------------------------------\n")
}

// main function.
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		underline()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		checkDNS("yahoo.com", "8.8.8.8")
		underline()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		checkDNS("google.com", "172.217.31.206")
		underline()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		checkDNS("apple.com", "4.4.4.4")
		underline()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		checkDNS("amazon.com", "1.1.1.1")
		underline()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		checkDNS("microsoft.com", "104.215.148.63")
		underline()
		wg.Done()
	}()

	// wait till the waitgroup counter is 0.
	wg.Wait()
}
