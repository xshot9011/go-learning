# Golang

- Based on version 1.18 **

## Go workspace

`gopls` support both Go module and GOPATH mode

### Module mode

- Single module

```bash
.
├── sub-module
│   └── auau.go
├── go.mod  # root module
└── main.go
```

When you open a parent directory that contains a module, it must contain only that module. Otherwise, you are working with multiple modules.

- Multiple module

Create a file named `go.work` that contains related modules, and then point your workspaces root directory to the directory containing the `go.work` file.

### Basic command

As an example, consider checking out this repository to the $WORK/tools directory. By creating a `go.work` file, we can simultaneously work on `golang.org/x/tools` and `golang.org/x/tools/gopls`.

```bash
cd <PATH>
go work init
go work use tools tools/gopls
```

## Golang

### Syntax

```go
package main

import "fmt"

func main() {
	// Main data type {string, int, bool, map, arrays}
	var message string = "initial"
	var message1 = "initial"
	message2 := "initial"
	fmt.Printf("%v\n%T\n%s\n", message, message1, message2)

	// Special variable "pointer"
	fmt.Printf("The message address is: %p, containe value: %s\n", &message, message)

	// Arrays and slices
	var books [50]string
	books[0] = "nano"
	books[49] = "The last emp"
	fmt.Printf("Books has len: %v, with %v\n", len(books), books)

	// var movies = []string{}
	// var movies []string
	movies := []string{}
	movies = append(movies, "King")
	movies = append(movies, "King2")
	movies = append(movies, "King3")
	fmt.Printf("Movies: %v\n", movies)

	// Loop
	for index, movie := range movies {
		fmt.Printf("Movie %v: %v\n", index, movie)
	}

	// If
	if movies[0] == "kings" && movies[1] == "Wow za" {
		// do something
	} else if movies[1] == "Kingalo" {
		// do something
	}

	// Switch
	switch movies[0] {
	case "kingra":
		// do something
	case "konga":
		// do something
	default:
		fmt.Printf("No match given\n")
	}

	// Function call
	response := do_nothing(movies)
	fmt.Printf("Response is: %v\n", response)
}

func do_nothing(slices []string) []string {
	return slices
}
```

Note: `If a function or variable is declared outside of a function, it is accessible from any other file in same package.`

### Variable Scope

- `Local` Variable defined within a function
- `Package` Outside-of-function variable
- `Global` Variable that defines a the outside function beginning with Capital letters, allowing it to be accessed by the entire package.

### Single Package

```
.
├── main.go  # Package main
├── connector  # Package main
└── validator  # Package main
```

### Multiple Packages

```
.
├── main.go  # Package main
└── module
    └── auah.go  # Package auah
```
