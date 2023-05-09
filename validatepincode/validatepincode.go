package validatepincode

import (
	"regexp"
	"unicode"
)

func IsValidPinCode(s string) bool {

	if len(s) < 6 || checkSequentialNumber(s) {
		return false
	}

	var prev rune
	var count, dubSet int
	for _, char := range s {
		if unicode.IsDigit(char) {
			if char == prev {
				count++
			} else {
				count = 1
			}
			if count > 1 {
				dubSet += 1
				count = 0
				if dubSet > 2 {
					return false
				}
			}
		} else {
			count = 0
		}

		prev = char
	}

	return true
}

func checkSequentialNumber(pinCode string) bool {
	re := regexp.MustCompile(`123|234|345|456|567|678|789|987|876|765|654|543|432|321`)
	return re.MatchString(pinCode)
}
