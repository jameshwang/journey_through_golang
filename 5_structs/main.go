package main

import (
	"fmt"
	// "time"
)

type geoPoint struct {
	latitude  float32
	longitude float32
}

type trafficLight struct {
	state    string
	location geoPoint
}

func (t *trafficLight) next() {
	if t.state == "Green" {
		t.state = "Yellow"
	} else if t.state == "Yellow" {
		t.state = "Red"
	} else {
		t.state = "Green"
	}
}

func main() {
	jacksonHyde := trafficLight{state: "Red"}
	fmt.Println(jacksonHyde.state)
	jacksonHyde.next()
	fmt.Println(jacksonHyde.state)
	jacksonHyde.next()
	fmt.Println(jacksonHyde.state)
}
