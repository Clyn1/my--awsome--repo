package main

import "fmt"

func DealAPackOfCards(deck []int) {
	if len(deck)%4 != 0 {
		fmt.Println("Cannot deal the cards evenly among 4 players.")
		return
	}

	cardsPerPlayer := len(deck) / 4

	fmt.Println("Player 1:")
	fmt.Println(deck[:cardsPerPlayer])

	fmt.Println("\nPlayer 2:")
	fmt.Println(deck[cardsPerPlayer : 2*cardsPerPlayer])

	fmt.Println("\nPlayer 3:")
	fmt.Println(deck[2*cardsPerPlayer : 3*cardsPerPlayer])

	fmt.Println("\nPlayer 4:")
	fmt.Println(deck[3*cardsPerPlayer:])
}

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
