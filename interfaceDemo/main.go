package main

import "fmt"

type bot interface {
	/* Any type that has a reciever function called
	   getGreeting() that returns a string is now
	   of type 'bot'. Basically, as long as something
	   satisfies the conditions of an interface, it can
	   be used in locations where that interface is used.
	*/
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//Assume there's logic here for calculating a greeting based on... something
	return "Hello"
}

func (spanishBot) getGreeting() string {
	//Assume custom logic for a spanish greeting
	return "Hola"
}
