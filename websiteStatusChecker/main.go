package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		printSiteStatus(link)
	}

}

func printSiteStatus(address string) {
	_, err := http.Get(address)

	if err != nil {
		fmt.Printf("%v connection error!!! \t\t%v\n", address, err)
	} else {
		fmt.Printf("%v connection successful!\n", address)
	}
	return
}
