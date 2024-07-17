package main

import "strings"

func calculateVocalChar(str string) int {
	// vocal aeiou
	total := 0
	lower := strings.ToLower(str)

	for _, v := range lower {
		if v == 'a' {
			total++
		} else if v == 'e' {
			total++
		} else if v == 'i' {
			total++
		} else if v == 'o' {
			total++
		} else if v == 'u' {
			total++
		}
	}

	return total
}
