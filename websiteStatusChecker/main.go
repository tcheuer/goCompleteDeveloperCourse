package main

import (
	"fmt"
	"net/http"
)

//Concurrency is not parallelism. Concurrency says that any go routine can
//execute in any order. Parallelism means that multiple go routines can execute
//at the same time.  Parallelism requires multiple cores.
func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://amazon.com",
	}

	//Remember channels can only work with single type.
	//The channel must be passed into the printSiteStatus function,
	//otherwise it will not be able to use it to communicate.
	c := make(chan string)

	for _, link := range links {
		//The go keyord is only used in front of function calls
		go printSiteStatus(link, c)
	}

	fmt.Println(<-c)

}

func printSiteStatus(address string, c chan string) {
	_, err := http.Get(address) //Blocking call!

	if err != nil {
		fmt.Printf("%v connection error!!! \t\t%v\n", address, err)
		c <- "May be down"
	} else {
		fmt.Printf("%v connection successful!\n", address)
		c <- "It is up"
	}
	return
}
