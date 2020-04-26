package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func multiple(x, y string) (string, string) {
	return x, y
}

func main() {
	var num1, num2 float64 = 2.3, 4.5
	//var num2 float64 = 4.5

	var a int = 34

	var b float64 = float64(a)

	c := a // c will be int
	x, y := "hey", "ho"
	fmt.Println(add(num1, num2))

	fmt.Println(multiple(x, y))
}
