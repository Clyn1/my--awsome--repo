package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func PrintStr(s string) {
	for _, j := range s {
		z01.PrintRune(j)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func base(mn int) {
	t := '0'
	for j := 0; j < mn; j++ {
		t++
	}
	z01.PrintRune(t)
}

func main() {
	partA := "x = "
	partB := ", y = "
	point := point{}
	setPoint(&point)
	PrintStr(partA)
	base(point.x / 10)
	base(point.x % 10)
	PrintStr(partB)
	base(point.y / 10)
	base(point.y % 10)
	z01.PrintRune('\n')
}
