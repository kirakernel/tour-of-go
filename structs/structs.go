package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type User struct {
	name string
	email string
	is_online bool
}

type Vtuber struct {
	user User
	channel string
	subscribers_number int
}

func main() {
	fmt.Println(Vertex {1, 2});
	user := User { "Nene", "nene@cute.uwu", true }
	fmt.Println(user)
	fmt.Println(Vtuber { user, "CuteGF", 300_000 })
}