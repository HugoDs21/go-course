package main

func main() {
	// var card string = "Ace of Spades"
	// var card = "Ace of Spades"
	// card := "Ace of Spades"
	//cards := newDeck()
	//cards.saveToFile("test")
	cards := newDeckFromFile("test")
	cards.shuffle()
	cards.print()
}
