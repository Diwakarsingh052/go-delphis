// package main files can be executed
package main

import "fmt"

var address = "random address"

// create a variable using shorthand here, and see what happens

// a:= 100 // it is not allowed here to declare var in global scope using shorthand

func main() {
	var name string = "dev"
	var age = 29
	marks := 100 // shorthand operator ( `:=` )
	// fmt package contains functionality related to printing and scanning
	fmt.Println("Hello Everyone", name, age, marks, address)
}

// go run 1-var.go // to run the program
