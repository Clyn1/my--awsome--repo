package piscine

import (
	"github.com/01-edu/z01"
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

func (piscine) {
	input := "1 2 3 4 5"
	intSlice, err := StringToIntSlice(input)
	if err != nil {
		z01.PrintRune("Error:", err)
		return
	}

	z01.PrintRune("Input String:", input)
	z01.PrintRune("Converted Int Slice:", intSlice)
}
