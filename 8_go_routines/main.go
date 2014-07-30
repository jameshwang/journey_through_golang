package main

import (
	"fmt"
	"time"
)

func sleeper(notify chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Second number: %d\n", i)
	}

	notify <- "done"
}

func sleep1Second(notify chan bool, id int) {
	time.Sleep(time.Second * 10)
	fmt.Printf("I'm done %d at %s\n", id, time.Now().String())
	notify <- true
}

func sleep2Second(notify chan bool, id int) {
	time.Sleep(time.Second * 20)
	fmt.Printf("I'm done %d at %s\n", id, time.Now().String())
	notify <- true
}

func main() {
	count := 10000
	notify := make(chan bool, count)
	for i := 0; i < count; i++ {
		if count%2 == 0 {
			go sleep2Second(notify, i)
		} else {
			go sleep1Second(notify, i)
		}
	}

	<-notify
}
