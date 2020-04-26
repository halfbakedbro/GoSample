package main

import "fmt"

func main() {
	x := 24
	y := &x // memmory address

	fmt.Println(y)
	fmt.Println(*y)

	*y = 12

	fmt.Println(x)
}
