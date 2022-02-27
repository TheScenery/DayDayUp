package tihuankonggelcof

import "strings"

func replaceSpace(s string) string {
	result := strings.Builder{}
	for _, ch := range s {
		if ch == ' ' {
			result.WriteString("%20")
		} else {
			result.WriteRune(ch)
		}
	}
	return result.String()
}
