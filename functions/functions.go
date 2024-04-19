package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(square(5))
	fmt.Println(square(10))
}

func square(x int) int {
	return  x * x
}
