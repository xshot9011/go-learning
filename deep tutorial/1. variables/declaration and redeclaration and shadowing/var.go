package main

import "fmt"

var number5 int = 115 // declare at package level, full term of declaration only

// declare at package level, with many variable
var (
	fullName string = "Mr. Alisabet Youlike"
	nickname string = "betty"
	age      int    = 64
)

func main() {
	var number int // declare the variable call "number" and it's type is int
	number = 112
	var number2 int = 113 // declare the variable call "number" and it's type is int
	number3 := 114        // auto assign value with type
	// note: you cannot redeclare the same name with exist vars name

	fmt.Println(number, number2, number3)
	fmt.Printf("number3 values is %v, type is %T\n", number3, number3) // %v is value, %T is type

	fmt.Printf("package var, number5 is %v\n", number5)
	var number5 int = 116 // do shadowing
	fmt.Printf("shadowing the variable, number5 is %v\n", number5)

	// go will detect none usage variable and raise exception
}
