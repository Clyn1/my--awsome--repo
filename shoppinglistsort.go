package piscine

import (
	"fmt"
	"sort"
)

func ShoppingListSort(items []string) []string {
	lengthMap := make(map[int][]string)

	for _, item := range items {
		length := len(item)
		lengthMap[length] = append(lengthMap[length], item)
	}

	var lengths []int
	for length := range lengthMap {
		lengths = append(lengths, length)
	}
	sort.Ints(lengths)

	var result []string
	for _, length := range lengths {
		result = append(result, lengthMap[length]...)
	}

	return result
}

func piscine() {
	shoppingList := []string{"eggs", "bread", "milk", "apple", "cheese"}
	sortedList := ShoppingListSort(shoppingList)

	fmt.Println("Original shopping list:", shoppingList)
	fmt.Println("Sorted shopping list:", sortedList)
}
