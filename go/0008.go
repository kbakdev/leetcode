package _go

import "strings"

func myAtoi(s string) int {
	var result int
	var sign int

	// prevent index out of range
	if len(s) == 0 {
		return 0
	}

	// make string to []rune
	runeString := []rune(s)

	// remove space
	for i := 0; i < len(runeString); i++ {
		if runeString[i] == ' ' {
			continue
		} else {
			runeString = runeString[i:]
			break
		}
	}

	// prevent index out of range
	if len(runeString) == 0 {
		return 0
	}

	// if runeString contains 2 elements or more
	if len(runeString) > 2 {
		if strings.ContainsRune(string(runeString[0]), '-') && strings.ContainsRune(string(runeString[1]), '+') {
			return 0
		}
	}

	// check number
	for i := 0; i < len(runeString); i++ {
		if runeString[i] >= '0' && runeString[i] <= '9' {
			result = result*10 + int(runeString[i]-'0')
		} else {
			break
		}
	}

	// check numbers overflow
	if result > 2147483647 {
		if sign == -1 {
			return -2147483648
		} else {
			return 2147483647
		}
	}

	// check overflow
	if result > 2147483647 {
		if sign == 1 {
			return 2147483647
		} else {
			return -2147483648
		}
	}

	// check sign
	if sign == -1 {
		result = -result
	} else if sign == +1 {
		result = +result
	}

	return result
}
