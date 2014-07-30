package main

import (
	// "errors"
	"fmt"
	"os"
)

func main() {

	// Arrays!
	// They have fixed lengths
	var a [3]int
	b := [2]string{"Hello", "World"}
	c := [...]string{"this", "can't", "be", "happening"}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Slices! Abstraction on top of Array
	d := []string{"a", "B"}
	fmt.Println(d)

	var e []string
	e = make([]string, 5, 15)

	fmt.Println(len(e))
	fmt.Println(cap(e))

	f := make([]string, len(e), (cap(e)+1)*2)
	copy(f, e)
	e = f

	fmt.Println(len(e))
	fmt.Println(cap(e))

	e = append(e, "3", "d")

	fmt.Println(e)

	g := map[string]int{"foo": 1, "bar": 2}

	for key, value := range g {
		fmt.Printf("key: %s and value: %b\n", key, value)
	}

	fmt.Println(g)

	// var h []byte
	// h := make([]byte, 64, 128)

	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	h, err := file.Read(b[0:64])

	fmt.Println(string(h))

}
