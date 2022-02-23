package stringtointegeratoi

import (
	"math"
)

func isNum(ch rune) bool {
	code := ch - '0'
	return code >= 0 && code <= 9
}

func myAtoi(s string) int {
	sign := 1
	num := 0
	start := false
	for _, ch := range s {
		if start && !isNum(ch) {
			break
		}
		if ch == ' ' {
			continue
		} else if ch == '-' {
			sign = -1
			start = true
			continue
		} else if ch == '+' {
			start = true
			continue
		}
		if !isNum(ch) {
			break
		}
		start = true
		result := num*10 + int(ch-'0')
		if sign < 0 && result*sign < math.MinInt32 {
			return math.MinInt32
		}
		if sign > 0 && result > math.MaxInt32 {
			return math.MaxInt32
		}
		num = result
	}
	return num * sign
}
