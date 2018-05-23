package main

import (
	"fmt"
	"net/http"
	"time"
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

	/* 	Simplistic example, only works with one value
	//Recieving anything from a channel is BLOCKING.
	//The program will wait here for another go-routine
	//to finish. After it finishes, the main routine ends,
	//eliminating any other running go routines.
	fmt.Println(<-c)
	*/

	/*	This example runs through and ensures that all go routines finish
		Based on the length of the slice.

		for i := 0; i < len(links); i++ {
			// <-c is a blocking call,
			//so it will pause execution until it recieves a message
			fmt.Println(<-c)
		}

	*/

	/* Standard infinite loop method
	//This is an intentionally infinite for loop.
	for {
		//This blocking request means that you will only run each time a link
		//is passed into the channel.
		go printSiteStatus(<-c, c)
	}
	*/

	/* This method is constant, but makes it difficult to control the timing of
	   without creating throttling.

	//This is a shorthand format of the previous loop. It says:
	//Wait for a value to appear in channel c, each time a value
	//is recieved run the loop with the value in variable l
	for l := range c {
		go printSiteStatus(l, c)
		//Putting this here pauses the main routine, which
		//means that a subgoroutine may finish and send a message, but
		//the message gets backed up and has to wat. You could have a
		//queue of 3 messages, then the 'website is up' message that
		//prints is alreay 15 seconds old when it prints.
		//time.Sleep(time.Second * 5)
	}
	*/

	//Function literal, used to wrap some code to cause it to run in the future
	//In C# these are called lambdas, in Java they're called anonymous functions
	for l := range c {

		//This function literal is referencing l, which is
		//in the outer scope.  As such, when printSiteStatus
		//is called, it's going to reference a memory location
		//that's being used by another go routine. It could
		//change the value while the function is executing,
		//resulting in unexpected behaviour.

		go func(address string) {
			time.Sleep(5 * time.Second)
			printSiteStatus(address, c)
		}(l)
		//These are needed to call the function. The value in 'l' is now
		//passed by value to the function literal by value, which keeps it
		//from referencing the original value.
	}

}

func printSiteStatus(address string, c chan string) {
	//Pausing here would be weird, because calling
	//the function implies that the link is going to
	//be checked immediately.
	//time.Sleep(time.Second * 5)
	_, err := http.Get(address) //Blocking call!

	if err != nil {
		fmt.Printf("%v connection error!!! \t\t%v\n", address, err)
		c <- address
	} else {
		fmt.Printf("%v connection successful!\n", address)
		c <- address
	}
	return
}
