package adddigits

//https://en.wikipedia.org/wiki/Digital_root
func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	r := num % 9
	if r == 0 {
		return 9
	}
	return r
}
