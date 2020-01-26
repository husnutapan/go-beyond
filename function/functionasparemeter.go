package main

import "fmt"

type convert func(int) string

func convertIntegerStr(function convert) string {
	return fmt.Sprintf("%q", function(111))
}

func main() {
	var result string
	myFunc := func(a int) string {
		return "myFunc"
	}
	result = convertIntegerStr(myFunc)
	fmt.Println(result)
}
