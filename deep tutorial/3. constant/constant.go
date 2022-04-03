package main

import "fmt"

/*
iota พบว่าเมื่อมันถูกใช้ใน constant expression แล้ว จะกำหนดค่าเริ่มต้นเป็น 0 เสมอ
	เมื่อเรียกซ้ำใน constant expression เดียวกัน มันจะเพิ่มขึ้นมา 1
	เพื่อความไม่งง ดังตัวอย่าง

ใช้ในการทำ enum
*/

const (
	a = iota
	b = iota
	c = iota
)

// is equal to
const (
	a2 = iota
	b2
	c2
)

const (
	a3 = iota + 5 // 5
	b3            // 6
	c3            // 7
)

// for unit
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// for permission
const (
	isAdmin = 1 << iota
	isHead
	isPo
	isMo
	isRead
	isWrite
	isEdit
)

func main() {
	const myConst int = 34 // declare constant
	const myConst2 = 68    // use compiler ability to infer the type of variable
	// const sinValue float64 = math.sin(xaxasxasx)  .. you cannot do this
	fmt.Printf("myConst values is %v\n", myConst)
	fmt.Printf("myConst2 values is %v\n", myConst2)

	// constant can be operated with the same non-const var type
	const a int = 15
	const a2 = 15
	var b int = 2
	var b2 int64 = 4
	// a + b is working fine
	// a + b2 is not working anymore
	fmt.Printf("a+b=%v, a+b2=%v", a+b, a+b2)
	// a2 + b is working fine
	// a2 + b2 is working fine, compiler infer the type
	fmt.Printf("a2+b=%v, a2+b2=%v", a2+b, a2+b2) // as Printf("", 15+b, 15+b2) work as implicit casting, like a symbol where a2 is there is 42

	var userRole byte = isAdmin | isRead | isWrite | isEdit
	// to check if has permission => userRole & isAdmin == isAdmin
	// to print permission => userRole
}
