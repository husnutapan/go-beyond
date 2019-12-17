package main

import "fmt"

type spanishBot struct{}
type englishBot struct{}

func main() {
	eb := englishBot{}
	printGreeting(eb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

//Go doesnt allow overload like java!!!
/*
func printGreeting(sp spanishBot) {
	fmt.Println(sp.getGreeting())
}
*/
func (englishBot) getGreeting() string {
	return "Hi There"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}
