package main

import "fmt"

// creating custom types
type money int

func main() {
	// we neeed to create the custom type variable using var keyword only
	var dollar money = 100

	fmt.Println(dollar)
	fmt.Printf("%v\n", dollar)
	fmt.Printf("%T\n", dollar)
}
