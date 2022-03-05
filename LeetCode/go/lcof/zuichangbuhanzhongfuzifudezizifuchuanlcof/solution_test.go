package zuichangbuhanzhongfuzifudezizifuchuanlcof

import "testing"

func TestFunc(t *testing.T) {
	s := "abcabcbb"
	r := lengthOfLongestSubstring(s)
	if r != 3 {
		t.Fatal(r)
	}

	s = " "
	r = lengthOfLongestSubstring(s)
	if r != 1 {
		t.Fatal(r)
	}
}
