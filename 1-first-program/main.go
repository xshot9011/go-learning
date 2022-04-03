package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var BANNER_MSG string = "HELLO WORLD"
	const CONSTANT string = "THIS IS CONSTANT"
	// CONSTANT = "REDEFINE CONSTANT"  // Due to the fact that this is a constant variable, it cannot be redefined.
	SUB_BANNER := "THIS IS SUB BANNER"
	SUB_BANNER = "REDEFINED SUB BANNER"

	fmt.Printf("[INFO] BANNER_MSG: %s\n", BANNER_MSG)
	fmt.Printf("[INFO] CONSTANT: %s\n", CONSTANT)
	fmt.Printf("[INFO] SUB_BANNER: %s\n", SUB_BANNER)
}
