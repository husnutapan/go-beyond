package main

import "fmt"

func main() {
	var a = 6
	var p = &a

	fmt.Println("a value is=", a)
	fmt.Println("a address is=", &a)
	fmt.Println("p value is=", p)
}
