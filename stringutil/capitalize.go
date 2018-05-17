package stringutil

import "errors"

// Capitalize returns the argument string in capital
func Capitalize(s string) string {
	ret := []rune(s)

	for i, char := range ret {
		temp, err := capitalize(char)

		if err != nil {
			continue
		}

		ret[i] = temp
	}

	return string(ret)
}

func capitalize(char rune) (rune, error) {
	switch {
	case char > 64 && char < 91:
		return char, nil
	case char > 96 && char < 123:
		return char - 32, nil
	default:
		return -1, errors.New("Can only capitalize characters")
	}
}
