package main

import "fmt"

func main() {
	cards := newDeck()
	cards.print()
	cards.upperCase()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
}
