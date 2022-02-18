package stringtointegeratoi

import "testing"

func TestFunc(t *testing.T) {
	s := "42"
	r := myAtoi(s)
	if r != 42 {
		t.Fatal(r)
	}

	s = "    -42"
	r = myAtoi(s)
	if r != -42 {
		t.Fatal(r)
	}

	s = "4193 with words"
	r = myAtoi(s)
	if r != 4193 {
		t.Fatal(r)
	}

	s = "words and 987"
	r = myAtoi(s)
	if r != 0 {
		t.Fatal(r)
	}

	s = "+1"
	r = myAtoi(s)
	if r != 1 {
		t.Fatal(r)
	}

	s = "+-12"
	r = myAtoi(s)
	if r != 0 {
		t.Fatal(r)
	}

	s = "0000-12"
	r = myAtoi(s)
	if r != 0 {
		t.Fatal(r)
	}

	s = "9223372036854775808"
	r = myAtoi(s)
	if r != 2147483647 {
		t.Fatal(r)
	}
}
