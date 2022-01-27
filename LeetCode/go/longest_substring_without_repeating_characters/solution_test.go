package longestsubstringwithoutrepeatingcharacters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	result := lengthOfLongestSubstring(s)
	if result != 3 {
		t.Fail()
	}

	s = "bbbbb"
	result = lengthOfLongestSubstring(s)
	if result != 1 {
		t.Fail()
	}

	s = "pwwkew"
	result = lengthOfLongestSubstring(s)
	if result != 3 {
		t.Fail()
	}

	s = ""
	result = lengthOfLongestSubstring(s)
	if result != 0 {
		t.Fail()
	}

	s = " "
	result = lengthOfLongestSubstring(s)
	if result != 1 {
		t.Fail()
	}
}
