package findtheclosestpalindrome

import "testing"

func TestFunc(t *testing.T) {
	n := "123"
	r := nearestPalindromic(n)
	if r != "121" {
		t.Fatal(r)
	}

	n = "1"
	r = nearestPalindromic(n)
	if r != "0" {
		t.Fatal(r)
	}
}
