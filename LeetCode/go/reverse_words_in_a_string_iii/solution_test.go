package reversewordsiii

import "testing"

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	o := "s'teL ekat edoCteeL tsetnoc"
	r := reverseWords(s)
	if r != o {
		t.Fatal(r)
	}
}
