package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"

	//var colors map[string]string

	//colors := make(map[string]string)

	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%s\t:\t%s\n", color, hex)
	}
}
