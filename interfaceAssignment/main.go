package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type shape interface {
	getArea() (float64, error)
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

func main() {
	tri := triangle{
		base:   3,
		height: 4,
	}

	sqr := square{
		side: 2,
	}

	printArea(tri)
	printArea(sqr)
}

//Returns the area of a triangle based on its
//base and height values. Retruns -1 and an
//error if unable to calculate the value.
func (t triangle) getArea() (float64, error) {

	if t.base < 0 || t.height < 0 {
		return -1, errors.New(fmt.Sprintf("could not calculate area due to negative value. Recieved: base: %v height: %v", t.base, t.height))
	}

	return 0.5 * t.base * t.height, nil
}

//Returns the area of a square
//Returns -1 and an error if unable to calculate the
//value.
func (s square) getArea() (float64, error) {

	if s.side < 0 {
		return -1, errors.New(fmt.Sprintf("could not calculate area due to negative value. Recieved: side: %v", s.side))
	}

	return s.side * s.side, nil
}

//Prints the area of anything that follows the
//shape interface to the console
func printArea(s shape) {
	shapeName := reflect.TypeOf(s)
	area, err := s.getArea()

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fmt.Printf("The %v has an area of : %v\n", shapeName.Name(), area)
}
