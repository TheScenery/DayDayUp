package zuoxuanzhuanzifuchuanlcof

func reverseLeftWords(s string, n int) string {
	l := s[0:n]
	r := s[n:]
	return r + l
}
