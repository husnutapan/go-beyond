package main

func main() {
	/*cards := newDeck()
	cards.print()
	cards.upperCase()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
	//save data to file with byte array
	cards.saveToFile("my_memory_cards.txt")

	fromFileDeck := newDeckFromFile("my_memory_cards.txt")
	fromFileDeck.print()*/

	cards := newDeck()
	cards.suffle()
	cards.print()

}
