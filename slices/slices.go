package main

import "fmt"

func main() {
	primes := [6]int{2,3,5,7,11,13}

	var s[]int = primes[1:4]
	fmt.Println(s)

	names := [3]string{"Nene", "Aia", "Nerissa"}

	var s2[]string = names[1:2]
	fmt.Println(s2)
}