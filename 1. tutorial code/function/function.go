package main

//package main

import "fmt"

func main() {
	var msg string
	var stage int
	msg, stage = getFullName("Big", "Tu")
	fmt.Printf("%s with stage %d", msg, stage)  // work as C
	fmt.Println("%s with stage %d", msg, stage) // work as python
}

func getFullName(firstName string, lastName string) (string, int) {
	return (firstName + " " + lastName), 0
}
