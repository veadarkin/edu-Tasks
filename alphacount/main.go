package main

import (
	"fmt"
)

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

// AlphaCount : подсчет букв латинского алфавита
func AlphaCount(s string) (count int) {
	for _, v := range s {
		if (v >= 97 && v <= 122) || // буквы в промежутке по таблице ascii
			(v >= 65 && v <= 90) {
			count++
		}
	}
	return
}
