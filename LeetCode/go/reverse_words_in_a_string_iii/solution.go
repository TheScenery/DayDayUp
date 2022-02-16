package reversewordsiii

import "strings"

func reverseWord(s string) string {
	sArray := []byte(s)
	n := len(sArray)
	l := 0
	r := n - 1
	for l < r {
		sArray[l], sArray[r] = sArray[r], sArray[l]
		l++
		r--
	}
	return string(sArray)
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = reverseWord(word)
	}
	return strings.Join(words, " ")
}
