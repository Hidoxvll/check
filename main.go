package main

import "fmt"

//
//import "fmt"
//
//type Animal interface {
//	Speack() string
//}
//type Dog struct {
//}
//
//func (Dog) Speack() string {
//	return "Woof!"
//}
//
//type Cat struct {
//}
//
//func (Cat) Speack() string {
//	return "Nice to meet you!"
//}
//
//func main() {
//	var a Animal
//
//	a = Dog{}
//	fmt.Printf("Dog voice: %s\n", a.Speack())
//	a = Cat{}
//	fmt.Printf("Cat voice: %s\n", a.Speack()

func chek(i interface{}) {
	fmt.Printf("Тип: %T, Значение: %v\n", i, i)
}

func main() {
	chek(123)
	chek("Salom")
	chek(12.21)
	chek(true)
	chek([]string{"Чёрный", "Белый", "Голубой"})

}
