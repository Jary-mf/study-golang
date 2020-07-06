// Package word provides utilities for word games
package word

import "unicode"

// IsPalindrome reports whether s reads the same forword and backward
// (Our first attempt)
func IsPalindrome(s string) bool {
	var letter []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letter = append(letter, unicode.ToLower(r))
		}
	}
	for i := range letter {
		if letter[i] != letter[len(letter)-1-i] {
			return false
		}
	}
	return true
}
