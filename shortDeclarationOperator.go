package main

import (
	"fmt"
)
var y = "hello world!"
func main() {
	x := 10

	fmt.Println(x, y)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T", z, z)
}