package piscine

import (
	"github.com/01-edu/z01"
)

type ByLength []string

func (s ByLength) Len() int           { return len(s) }
func (s ByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByLength) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

func ShoppingListSort(items []string) []string {
	sortedItems := make([]string, len(items))
	copy(sortedItems, items)
	Sort(ByLength(sortedItems))
	return sortedItems
}

func main() {
	shoppingList := []string{"eggs", "bread", "milk", "apple", "cheese"}
	sortedList := ShoppingListSort(shoppingList)

	for _, item := range sortedList {
		for _, char := range item {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
