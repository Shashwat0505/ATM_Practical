package functions

import "unicode"

func IsValid(n string) bool {

	for _, val := range n {
		if unicode.IsLetter(val) || unicode.IsSymbol(val) || unicode.IsSpace(val) || unicode.IsPunct(val) || unicode.IsPunct(val) {
			return false
		}
	}
	return true
}
