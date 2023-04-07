package functions

import "strings"

func CheckforDot(s string) bool {
	if strings.Contains(s, ".") {
		return true
	}
	return false
}
