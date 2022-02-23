package reverseinteger

import (
	"strconv"
)

func reverseString(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func reverse(x int) int {
	s := strconv.FormatInt(int64(x), 10)
	if s[0] == '-' {
		s = "-" + reverseString(s[1:])
	} else {
		s = reverseString(s)
	}
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int(v)
}
