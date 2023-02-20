package utils

import "strings"

func CheckHasVowels(in string, vowels string) bool {
	for _, d := range vowels {
		if strings.Contains(in, string(d)) {
			return true
		}
	}
	return false
}
