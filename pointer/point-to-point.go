package main

import "fmt"

func main() {
	var a = 5
	var p = &a
	var pp = &p

	fmt.Println("a =", a)
	fmt.Println("address of a=", &a)

	fmt.Println("p =", p)
	fmt.Println("address of p = ", &p)

	fmt.Println("*pp = ", pp)
	fmt.Println("**pp = ", **pp)

}
