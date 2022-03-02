package findtheclosestpalindrome

import (
	"math"
	"strconv"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func nearestPalindromic(n string) string {
	l := len(n)
	candidates := []int{int(math.Pow10(l-1) - 1), int(math.Pow10(l) + 1)}
	half := n[0 : (l+1)/2]
	halfInt, _ := strconv.Atoi(half)
	cases := []int{halfInt, halfInt + 1, halfInt - 1}
	for _, c := range cases {
		result := c
		if l%2 != 0 {
			c = c / 10
		}
		for c > 0 {
			result = result*10 + c%10
			c = c / 10
		}
		candidates = append(candidates, result)
	}

	r := -1
	num, _ := strconv.Atoi(n)
	for _, c := range candidates {
		if c == num {
			continue
		}
		if r == -1 {
			r = c
		} else {
			if abs(c-num) < abs(r-num) {
				r = c
			} else if abs(c-num) == abs(r-num) && c < r {
				r = c
			}
		}
	}

	return strconv.Itoa(r)
}
