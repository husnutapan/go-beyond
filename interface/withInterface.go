package main

import "fmt"

type bot interface {
	getGreeting1() string
}

type spanishBot1 struct{}
type englishBot1 struct{}

func main() {
	eb := englishBot1{}
	sb := spanishBot1{}
	printGreeting1(eb)
	printGreeting1(sb)

}

func printGreeting1(b bot) {
	fmt.Println(b.getGreeting1())
}

func (englishBot1) getGreeting1() string {
	return "Hi There"
}

func (spanishBot1) getGreeting1() string {
	return "Hola"
}
