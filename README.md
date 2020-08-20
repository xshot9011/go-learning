# Go

เพื่อการเรียนรู้โครงสร้างและไวยกรณ์ภาษา c

### run เพื่อดูผลเฉยๆ

```bash
go run [file_name.go]
```

### ต้องการตัว executable file

```bash
go build [file_name.go]
file_name.exe
```

### การทำให้โค๊ดอยู่ใน convention เดียวกัน

```bash
gofmt -w [file_name.go]
```

-w หมายถึงการเขียนผลลัพธ์ทับไฟล์เดิม

### การ auto import package

```bash
go get golang.org/x/tools/cmd/goimports  # to install
goimports -w [file_name.go]
```

## Control Statements

### if statement

```go
if condition{
    // statement
}else{
    // statement
}
```

### for-loop

```go
for i := 1; i <= 10; i++{
    // statement
}
```

### switch

```go
switch var {
  case value: // Statement
  case value: // Statement
  default: // Statement
}
```

## Variable Type

|Type|default value|
|--|--|
|bool|false|
|int|0|
|float|0.0|
|string|""|
|function|nil|

### declaration and assign: normal

```go
var [name] [type]  // after declare you can assign the value
[name] = value
```

เช่น
```go
var msg string
msg = "Hello, Go"
```

### assign: Type inference

```go
msg := "Hello, Go"  // type of msg is string
```

### Array

syntax:

```go
var [name] [[size]][type]  // declare with fixed size
var [name] [][type]  // declare with non-fixed size
```

#### assign: normal

```go
var msg [3]string
msg[0] = "Hello"
msg[1] = ","
msg[2] = "Go"
// or
var msg [3]string{"Hello", ",", "Go"}
// or
var msg []string  // non-fixed size
msg = append(msg, "Hello")
msg = append(msg, ",")
msg = append(msg, "Go")
```

#### assign: Type Inferece

```go
msg := [3]string{"Hello", ",", "Go"}
// or
msg := []string{"Hello", ",", "Go"}
```

#### Loop through Array

```go
for index, name := range name{
    fmt.PrintLn(index, name)
}
```

### Function

```go
package main

import "fmt"

func main() {
	var msg string
	var stage int
	msg, stage = getFullName("Big", "Tu")
	fmt.Printf("%s with stage %d", msg, stage)  // work as C
	fmt.Println("%s with stage %d", msg, stage) // work as python
}

// declare parameters as [name] [type]
func getFullName(firstName string, lastName string) (string, int) {
	return (firstName + " " + lastName), 0  // one function can return multiple value as declaration
}

```

### Struct

```go
type human struct {
  // Unexported struct fields are invisible to the JSON package.
  // Export a field by starting it with an uppercase letter.
  name string // Unexported
  agr int // Unexported
  Name string `json:"name"` // Export 
}
```

The underlying reason for this requirement is that the JSON package uses reflect to inspect struct fields. Since reflect doesn't allow access to unexported struct fields, the JSON package can't see their value.

เวลาใช้งานกับการเขียนเว็ปถ้าเราไม่ทำให้เป็น Exported field json จะมองไม่เห็น อันนี้อารมณ์เกือบเหมือนกับ serializer ใน django ที่เรา set attr ของ field นั้นได้ ส่วนนี้ต้องศึกษาเพิ่มเติม

usage

```go
type human struct {
	Name string
	Age  int
}

func (h human) printInfo2() {  // only human type can access this method
	fmt.Println(h.Name, h.Age)
}

func main() {
	person := human{Name: "O", Age: 15}
	person.Age = 68
	printInfo1(person)
	person.printInfo2()
}

func printInfo1(h human) {
	fmt.Println(h.Name, h.Age)
}
```

### Interfaces และ Duck Typing

```go
package main

import "fmt"

type human struct {
    name string
    age  int
}

type parrot struct {
    name string
    age  int
}

type talker interface {
    talk() // ถ้าเป็น talker ก็จะสามารถพูดได้
}
// กำหนดความสามารถผูกกับ struct นั้นๆ
func (h human) talk() {
    fmt.Println("Human - I'm talking.")
}

func (p parrot) talk() {
    fmt.Println("Parrot  - I'm talking.")
}

func main() {
    talkers := [2]talker{  // ประกาศ interface
        human{name: "Somchai", age: 23},
        parrot{name: "Dum", age: 2},
    }
    for _, talker := range talkers {
        talker.talk()
    }
    // ผลลัพธ์
    // Human - I'm talking.
    // Parrot  - I'm talking.
}
```

จากตัวอย่างพบว่าทั้งนกแก้วและคนต่างเป็นนักพูด เราจึงประกาศ interface ชื่อ talker ขึ้นมา การที่จะทำให้ Go รู้ว่าคนและนกแก้วเป็นนักพูดนั้น เราไม่ต้อง implements interface แบบภาษาอื่น Go ถือหลักการของ Duck Typing นั่นคือ ถ้าคุณร้องก๊าบๆและเดินเหมือนเป็ด คุณก็คือเป็ด หาก struct คุณมีเมธอด talk คุณก็คือ taker นั่นเอง!

### Concurrency และ Parallelism

การที่เราจำทำ concurrency ด้วย go นั้นง่ายมากและประสิทธิภาพดีเช่นกัน ด้วนตัว Goroutines ของ go

concurrency := การทำงานที่อิสระแยกจากกัน 

เมื่อทุกอย่างเป็นอิสระต่อกันจะเริ่มทำที่ไหนก่อนก็ไม่มีปัญหา ก็แบ่งให้ cpu แต่ละ core เอางานไปทำเลย การทำงานแบบขนานนี้เราเรียกว่า parallelism

#### Goroutines

เราสามารถสร้างการทำงานแบบ Concurrency ได้ด้วยการใช้ Goroutines เพียงแค่เติม "go" เข้าไปหน้าฟังก์ชั่นทุกอย่างก็เป็นที่เรียบร้อย

แต่ตอนนี้จะมีปัญหาเมื่อมันทำงานไปสุด main แล้วอะ แต่อันที่เรา go ไปมันยังทำงานไม่เสร็จเลย

สุด main ก็คือสุดการทำงาน ดังนั้นเราต้องบอก go ว่าคอยการทำงานของ Goroutines ให้เสร็จก่อน

วิธีที่เราจะใช้นั้นก็คือ WaitGroup ภายใต้ package sync

```go
import "sync"
// โปรดสังเกต เราต้องรับพอยเตอร์ของ sync.WaitGroup เข้ามาด้วย
func searchFromFolder(keyword string, folder string, wg *sync.WaitGroup) {
  // ทำการค้นหา
  // เมื่อค้นหาเสร็จ ต้องแจ้งให้ WaitGroup ทราบว่าเราทำงานเสร็จแล้ว
  // WaitGroup จะได้นับถูกว่าเหลือ Goroutines ที่ต้องรออีกกี่ตัว
  wg.Done()
}

func search(keyword string) {
  folders := [3]string{"Document", "Image", "Library"}
  var wg sync.WaitGroup
  wg.Add(len(folders))  // บอก go ว่ารอการทำงานของ Goroutines ทั้งหมดกี่ตัว
  for _, folder := range folders {
    // เราต้องส่ง reference ของ wg ไปด้วย เพื่อที่จะสั่ง Done
    go searchFromFolder(keyword, folder, &wg)  // ทำ concurrency
  }
  wg.Wait() // และเพื่อป้องกันไม่ให้ Go หยุดการทำงานไปในทันที เราจึงต้อง Wait จนกว่า Goroutines จะทำงานเสร็จหมด
}

func main() {
  search("dog")
}
```

### pointer

```go
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

```

ใช้งานเหมือนในภาษา c เลยยย
