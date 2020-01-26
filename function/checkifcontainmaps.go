package main

import "fmt"

func main() {

	dictionary := map[string]int{"1": 1, "2": 2, "3": 3}

	value, ok := dictionary["1"]

	if ok {
		fmt.Println(value)
	} else {
		fmt.Println(value)
	}
}
