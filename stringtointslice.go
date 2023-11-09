package piscine

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToIntSlice(input string) ([]int, error) {
	strNums := strings.Fields(input)

	var intSlice []int
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			return nil, err
		}
		intSlice = append(intSlice, num)
	}

	return intSlice, nil
}

func main() {
	input := "1 2 3 4 5"
	intSlice, err := StringToIntSlice(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Input String:", input)
	fmt.Println("Converted Int Slice:", intSlice)
}
