package zigzagconversion

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rowsCount := numRows
	if len(s) < rowsCount {
		rowsCount = len(s)
	}
	rows := make([]string, rowsCount)
	r := 0
	diff := -1
	for _, ch := range s {
		rows[r] += string(ch)
		if r == 0 || r == numRows-1 {
			diff *= -1
		}
		r += diff
	}
	return strings.Join(rows, "")
}
