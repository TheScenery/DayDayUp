package validparentheses

import "testing"

func TestFunc(t *testing.T) {
	s := "()"
	r := isValid(s)
	if r != true {
		t.Fatal(r)
	}

	s = "()[]{}"
	r = isValid(s)
	if r != true {
		t.Fatal(r)
	}

	s = "(]"
	r = isValid(s)
	if r != false {
		t.Fatal(r)
	}

	s = "([)]"
	r = isValid(s)
	if r != false {
		t.Fatal(r)
	}
}
