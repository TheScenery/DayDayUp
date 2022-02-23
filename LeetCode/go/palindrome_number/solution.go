package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	newNumber := 0
	v := x
	for v > 0 {
		newNumber = newNumber*10 + v%10
		v = v / 10
	}

	return newNumber == x
}
