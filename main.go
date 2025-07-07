package main

import "fmt"

type Animal interface {
	Speack() string
}
type Dog struct{}

func (Dog) Speack() string {
	return "Woof!"
}

type Cat struct{}

func (Cat) Speack() string {
	return "Nice to meet you!"
}

func main() {
	var a Animal
	a = Dog{}
	fmt.Printf("Dog voice: %s\n", a.Speack())
	a = Cat{}
	fmt.Printf("Cat voice: %s\n", a.Speack())
}
