package stringutil

import (
	"strings"
)

// Reverse returns it's argument string reversed.
// This is my implementation.
func Reverse(s string) string {
	var ret []string
	for i := len(s) - 1; i >= 0; i-- {
		ret = append(ret, string(s[i]))
	}

	return strings.Join(ret, "") // Used strings :(
}

// ReverseNew returns it's argument string reversed
// by converting it to rune.
func ReverseNew(s string) string {
	ret := []rune(s)
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return string(ret)
}
