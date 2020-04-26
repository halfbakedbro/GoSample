package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, val int) {
	defer wg.Done()
	c <- val * 3
}

func main() {
	val := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(val, i)
	}

	wg.Wait()
	close(val)
	for item := range val {
		fmt.Println(item)
	}
}
