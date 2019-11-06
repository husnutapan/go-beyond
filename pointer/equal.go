package main

import "fmt"

func main() {
	var a = 10
	var p1 = &a
	var p2 = &a

	if p1 == p2 {
		fmt.Println("Both pointers")
	}
}
