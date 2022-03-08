package reversewordsinastring

import "strings"

func reverseWords(s string) string {
	n := len(s)
	i := n - 1
	j := i
	r := &strings.Builder{}
	for i >= 0 && j >= 0 {
		for j >= 0 && s[j] == ' ' {
			j--
			i--
		}
		for i >= 0 && s[i] != ' ' {
			i--
		}
		r.WriteString(s[i+1 : j+1])
		if i >= 0 {
			r.WriteByte(' ')
		}
		j = i
	}
	return strings.TrimRight(r.String(), " ")
}
