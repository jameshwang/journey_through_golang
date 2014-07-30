package main

import (
	"code.google.com/p/go-uuid/uuid"
	"fmt"
	"time"
)

const (
	s string = "a constant string"
)

func main() {

	// Playing around with variable and type assignments
	var a string = "Hello there"
	var value int = 1234
	var value1 float32 = 3234.4
	b := "My name is James"

	fmt.Println(a)
	fmt.Println(value)
	fmt.Println(value1)
	fmt.Println(b)

	// Control Flow statements
	if 1 == 2 {
		fmt.Println("This can't be right")
	} else {
		fmt.Println("this should be right")
	}

	words := 3

	switch words {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("a was not assigned")
	}

	for i := 0; i < len(a); i++ {
		fmt.Printf("%c\n", a[i])
	}

	fmt.Printf("The time is: %s\n", time.Now())

	fmt.Println(uuid.New())
	fmt.Printf("%s", uuid.Version())

}
