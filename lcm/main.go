package main

import (
	"fmt"
)

func main() {
	fmt.Println(Lcm(2, 7))
	fmt.Println(Lcm(-3, 124))
}

func Lcd(a, b int) int {
	if a == b {
		return a
	}
	d := a - b
	if d < 0 {
		d = -d
		return Lcd(a, d)
	} else {
		return Lcd(b, d)
	}
}

func Lcm(a, b int) int {
	if a <= 0 || b <= 0 {
		return 0
	}
	return a * b / Lcd(a, b)
}
