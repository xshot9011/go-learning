package main

import "fmt"

func main() {
	var is_real bool // initial value is false
	fmt.Printf("is_real value is %v type is %T\n", is_real, is_real)

	var number int   // initial value is 0
	var number2 uint // initial value is 0
	fmt.Printf("number value is %v type is %T\n", number, number)
	fmt.Printf("number2 value is %v type is %T\n", number2, number2)

	a := 10
	b := 3
	fmt.Println(a+b, a-b, a*b, a/b, a%b) // int divide with it not cause the float

	var c int = 10
	var d int16 = 20
	// fmt.Println(c + d) // invalid operation: a + b (mismatched types int and int16)
	fmt.Println(c + int(d)) // Even if int and int16 is very similar, but you need to do explicit casting

	// bit operation
	// a = 1010
	// b = 0011
	fmt.Println(a&b, a|b, a^b, a&^b)
	fmt.Println(a>>2, a<<2) // 1010 > 0101 > 0010, 1010 < 10100 < 101000

	// floating point
	f := 3.14
	f = 13.3e10
	fmt.Println("f is", f)

	g := 10.0
	h := 3.0
	fmt.Println(g+h, g-h, g*h, g/h) // int divide with it not cause the float

	// complex
	var com complex64 = 1 + 2i
	var com2 complex64 = 2 + 3i
	var com3 complex64 = complex(3, 4) // 3 + 4i
	fmt.Printf("%v, %T\n", com3, com3)
	fmt.Println(com+com2, com-com2, com*com2, com/com2)
	fmt.Println(real(com+com2), imag(com+com2))

	//string
	s := "this is the message to my lovely dog"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v %v, %T\n", s[2], string(s[2]), s[2]) // string is just sequence of bytes

	b_s := []byte(s) // convert string to collection of byte
	fmt.Printf("%v, %T\n", b_s, b_s)

	//rune
	//Rune is a Type. It occupies 32bit and is meant to represent a Unicode CodePoint.
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
