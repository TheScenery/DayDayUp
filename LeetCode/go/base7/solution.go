package base7

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	s := make([]byte, 0)
	v := num
	if num < 0 {
		v *= -1
	}

	for v > 0 {
		s = append(s, '0'+byte(v%7))
		v /= 7
	}
	if num < 0 {
		s = append(s, '-')
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
