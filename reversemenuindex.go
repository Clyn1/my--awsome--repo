package piscine

import "fmt"

func ReverseMenuIndex(menu []string) []string {
	reversed := make([]string, len(menu))
	for i := 0; i < len(reversed); i++ {
		reversed[i] = menu[len(menu)-1-i]
	}
	return reversed
}

func main() {
	menu := []string{"pizza", "pasta", "salad", "soup"}
	fmt.Println("Original menu:", menu)
	reversed := ReverseMenuIndex(menu)
	fmt.Println("Reversed menu:", reversed)
}
