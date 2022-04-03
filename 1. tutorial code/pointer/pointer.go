package main

import "fmt"

func main() {
	num := 0
	increase1(num)
	fmt.Printf("before> num : %d\n", num)
	increase2(&num)
	fmt.Printf("after> num : %d", num)
}

func increase1(number int) {
	number++
}

func increase2(number *int) {
	*number++
}
