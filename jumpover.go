package piscine

import (
	"github.com/01-edu/z01"
)

func JumpOver(s string) string {
	if len(s) == 0 || len(s) < 3 {
		return "\n"
	}

	var result string
	for i := 2; i < len(s); i += 3 {
		result += string(s[i])
	}
	result += "\n"
	return result
}

func piscine() {
	for _, i := range JumpOver("Hello, world") {
		z01.PrintRune(i)
	}
}
