package main

import "fmt"

func main() {
	a, b := 10, "dev" // create and assign the values

	//var a int = 100 // cant do this // a redeclared in this block // not allowed

	b, c := "shiv", 100 // for b var  it would update and assign the value, but for c it would create and assign
	fmt.Println(a, b, c)
}
