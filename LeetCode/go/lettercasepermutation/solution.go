package lettercasepermutation

import (
	"strings"
	"unicode"
)

func letterCasePermutation(s string) []string {
	ans := []string{""}
	for _, ch := range s {
		n := len(ans)
		newAnswers := ans
		if unicode.IsLetter(rune(ch)) {
			newAnswers = make([]string, 2*n)
			for i, j := 0, 0; i < n; i, j = i+1, j+2 {
				newAnswers[j] = ans[i] + strings.ToLower(string(ch))
				newAnswers[j+1] = ans[i] + strings.ToUpper(string(ch))
			}
		} else {
			for i := 0; i < n; i++ {
				newAnswers[i] = ans[i] + string(ch)
			}
		}
		ans = newAnswers
	}
	return ans
}
