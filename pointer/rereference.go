package main

import "fmt"

func main() {

	var a = 10
	var x = &a

	fmt.Println("a = ", a)
	fmt.Println("x = ", x)
	fmt.Println("*x = ", *x)

	*x = 20
	fmt.Println("a new value = ", a)

}
