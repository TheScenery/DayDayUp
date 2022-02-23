package palindromenumber

import "testing"

func TestFunc(t *testing.T) {
	x := 121
	r := isPalindrome(x)
	if r != true {
		t.Fatal(r)
	}

	x = -121
	r = isPalindrome(x)
	if r != false {
		t.Fatal(r)
	}

	x = 10
	r = isPalindrome(x)
	if r != false {
		t.Fatal(r)
	}
}
