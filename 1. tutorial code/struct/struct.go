package main

import "fmt"

type human struct {
	name string
	age  int
}

func (h human) printInfo2() {
	fmt.Println(h.name, h.age)
}

func main() {
	person := human{name: "O", age: 15}
	person.age = 68
	printInfo1(person)
	person.printInfo2()
}

func printInfo1(h human) {
	fmt.Println(h.name, h.age)
}
