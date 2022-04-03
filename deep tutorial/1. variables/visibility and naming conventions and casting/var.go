package main

import (
	"fmt"
	"strconv"
)

/*
If name start with "uppercase", it will be exported var.
meant that: It can be accessed by outside of package
else: only in the same pakage
*/

var i int = 112
var I int = 112

func main() {
	// explicit casting
	var intNumber int = 112
	fmt.Printf("%v, %T\nAfter casting\n", intNumber, intNumber)
	var flaotNumber float32
	flaotNumber = float32(intNumber)
	fmt.Printf("%v, %T\n", flaotNumber, flaotNumber)

	/*
		var f float32 = 112.112
		var i int32
		i = f  // this would raise run time error, go cannot losing their info through the convertion
		i = int(f) // this will work
	*/

	// convert interger to the string

	var number int = 65
	var str string
	str = string(number)       // return 'A'
	str = strconv.Itoa(number) // convert number to "number"
	fmt.Printf("\n\n%v, %T", str, str)

}
