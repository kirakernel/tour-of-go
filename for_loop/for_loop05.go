package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("\n")
		for j := 0; j <= 10; j++ {
			fmt.Printf("%v x %v = %v.\n", i, j, i*j)
		}
	}
}
