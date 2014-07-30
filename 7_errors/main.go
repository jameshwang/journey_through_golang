package main

import (
	"errors"
	"fmt"
)

func doIGetErrors() (int, error) {
	return 1, errors.New("Hello")
}

func main() {
	// err := errors.New("This is an error")
	value, err := doIGetErrors()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}
