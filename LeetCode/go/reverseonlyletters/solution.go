package reverseonlyletters

import "unicode"

func reverseOnlyLetters(s string) string {
	chArr := []rune(s)
	i, j := 0, len(s)-1
	for i < j {
		if !unicode.IsLetter(chArr[i]) {
			i++
		}
		if !unicode.IsLetter(chArr[j]) {
			j--
		}
		if unicode.IsLetter(chArr[i]) && unicode.IsLetter(chArr[j]) {
			chArr[i], chArr[j] = chArr[j], chArr[i]
			i++
			j--
		}
	}
	return string(chArr)
}
