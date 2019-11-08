package main

import "fmt"

func main() {
	//given err
	/*numbers := make([]int, 5)
	numbers[6] = 10*/

	numbers := make([]int, 5)
	defer avoidException()
	numbers[6] = 10
	fmt.Println("Main func completed...")
	numbers[7] = 123

}

func avoidException() {
	if err := recover(); err != nil {
		fmt.Println("Avoiding given exception")
	}

}
