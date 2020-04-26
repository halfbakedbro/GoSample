package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Alas Recovered", r)
	}
	wg.Done()
}

func says(s string) {
	defer cleanup()
	for i := 0; i < 4; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)

		if i == 2 {
			panic("Ohh god No!!")
		}
	}
}

func main() {
	wg.Add(1)
	go says("Hey")
	wg.Add(1)
	go says("there")
	wg.Wait()
}
