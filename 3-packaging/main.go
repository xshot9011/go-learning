package main

import (
	"fmt"
	"packaging/otherpackage"
)

func main() {
	fmt.Printf("Hello World\n")
	doNothing()
	otherpackage.DoNothing()
}

func doNothing() {
	fmt.Printf("Do nothing function invoked")
}
