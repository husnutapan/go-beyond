package main

import "fmt"

func main() {
	doSth()
}

//defer work like lifo
func doSth() {
	defer subProcess("x")
	defer subProcess("y")
	defer subProcess("z")

	numbers := []int{1, 5, 6, 7, 4}

	for _, n := range numbers {
		fmt.Println(n)
	}

}

func subProcess(s string) {
	fmt.Println(s)
}
