package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func swap_b(x, y int) (int, int) {
	return y, x
}

func swap_c(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	c, d := swap_b(10, 5)
	fmt.Println(c, d)

	e, f := swap_c("love", "Nene")
	fmt.Println(e, f)
}