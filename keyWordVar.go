package main

import (
	"fmt"
)

var y = 20

func main(){
	x := 10
	anything(x)
}

func anything(x int){
	fmt.Println(y)
	fmt.Println(x)
}