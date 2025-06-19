package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Greet() {
	fmt.Println("Здравствуй,", u.Name)
}
func main() {
	users := []User{
		{Name: "Vasya", Age: 12},
		{Name: "Vanya", Age: 12},
		{Name: "serega", Age: 13},
		{Name: "dima", Age: 14},
		{Name: "lena", Age: 15},
		{Name: "katya", Age: 15},
		{Name: "sveta", Age: 33},
		{Name: "kyra", Age: 111},
	}
	for _, name := range users {
		name.Greet()
	}
}
