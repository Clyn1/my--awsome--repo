package piscine

import (
	"fmt"
)

func DealAPackOfCards(deck []int) {
	if len(deck)%4 != 0 {
		fmt.Println("Cannot deal the cards evenly among 4 players.")
		return
	}

	cardsPerPlayer := len(deck) / 4

	for i := 0; i < 4; i++ {
		player := i + 1
		playerCards := deck[i*cardsPerPlayer : (i+1)*cardsPerPlayer]
		fmt.Printf("Player %d:\n", player)
		fmt.Println(playerCards)
		fmt.Println()
	}
}

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
