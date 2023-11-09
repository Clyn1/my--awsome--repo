package piscine

import (
	"fmt"
	"sort"
)

func PodiumPosition(competitors [][]string) [][]string {
	podium := make([][]string, 3)
	for i := 0; i < 3; i++ {
		podium[i] = competitors[i]
	}
	sort.Slice(podium, func(i, j int) bool {
		return podium[i][1] < podium[j][1]
	})
	return podium
}

func main() {
	competitors := [][]string{
		{"Lewis Hamilton", "1"},
		{"Max Verstappen", "2"},
		{"Sergio Perez", "3"},
		{"Valtteri Bottas", "4"},
		{"Charles Leclerc", "5"},
		{"Lando Norris", "6"},
		{"Carlos Sainz", "7"},
		{"Daniel Ricciardo", "8"},
		{"Sebastian Vettel", "9"},
		{"Fernando Alonso", "10"},
	}
	fmt.Println("Original competitors:", competitors)
	podium := PodiumPosition(competitors)
	fmt.Println("Podium positions:", podium)
}
