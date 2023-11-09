package piscine

import (
	"fmt"
)

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reversedMenu := make([]string, length)

	for i := 0; i < length; i++ {
		reversedMenu[i] = menu[length-1-i]
	}

	return reversedMenu
}

func piscine() {
	menu := []string{"DishA", "DishB", "DishC", "DishD"}
	reversedMenu := ReverseMenuIndex(menu)

	fmt.Println("Original Menu:", menu)
	fmt.Println("Reversed Menu:", reversedMenu)
}
