package piscine

import (
	"fmt"
)

type menu struct {
	burger  int
	chips   int
	nuggets int
}

func FoodDeliveryTime(order string) (int, error) {
	m := menu{15, 10, 12}
	switch order {
	case "burger":
		return m.burger, nil
	case "chips":
		return m.chips, nil
	case "nuggets":
		return m.nuggets, nil
	default:
		return 0, fmt.Errorf("404")
	}
}

func main() {
	order := "burger"
	time, err := FoodDeliveryTime(order)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The time it takes to prepare %s is %d minutes.\n", order, time)
	}
}
