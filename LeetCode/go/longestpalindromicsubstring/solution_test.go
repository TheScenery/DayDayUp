package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	result := longestPalindrome(s)
	if result != "aba" {
		t.Fatal(result)
	}

	s = "cbbd"
	result = longestPalindrome(s)
	if result != "bb" {
		t.Fatal(result)
	}

	s = "a"
	result = longestPalindrome(s)
	if result != "a" {
		t.Fatal(result)
	}

	s = "ac"
	result = longestPalindrome(s)
	if result != "a" {
		t.Fatal(result)
	}

	s = "bb"
	result = longestPalindrome(s)
	if result != "bb" {
		t.Fatal(result)
	}
}
