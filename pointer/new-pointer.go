package main

import "fmt"

func main() {
	ptr := new(string)
	*ptr = "5"

	fmt.Println("Pointer=", ptr)
	fmt.Println("Pointer Value=", *ptr)
}
