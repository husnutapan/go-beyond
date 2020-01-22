package main

import "fmt"

func sumOfTwoNumber(number1 int, number2 int, total chan int) {
	sum := number1 + number2
	total <- sum
}

func substractOfTwoNumber(number1 int, number2 int, total1 chan int) {
	sum := number1 - number2
	total1 <- sum
}

func main() {
	total := make(chan int)
	total1 := make(chan int)
	go sumOfTwoNumber(100, 200, total)
	go substractOfTwoNumber(100, 200, total1)

	sum, subtract := <-total, <-total1
	fmt.Printf("Result sum is:%d, substract is:%d", sum, subtract)

}
