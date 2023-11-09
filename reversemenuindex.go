package piscine

import "strings"

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reversedMenu := make([]string, length)

	for i := 0; i < length; i++ {
		reversedMenu[i] = menu[length-1-i]
	}

	return reversedMenu
}

func PrintMenu(menu []string) {
	joinedMenu := strings.Join(menu, " ")
	println(joinedMenu)
}

func main() {
	menu := []string{"DishA", "DishB", "DishC", "DishD"}
	reversedMenu := ReverseMenuIndex(menu)

	PrintMenu(menu)
	PrintMenu(reversedMenu)
}
