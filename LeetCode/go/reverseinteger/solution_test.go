package reverseinteger

import "testing"

func TestFunc(t *testing.T) {
	x := 123
	r := reverse(x)
	if r != 321 {
		t.Fatal(r)
	}

	x = -123
	r = reverse(x)
	if r != -321 {
		t.Fatal(r)
	}

	x = 120
	r = reverse(x)
	if r != 21 {
		t.Fatal(r)
	}

	x = 0
	r = reverse(x)
	if r != 0 {
		t.Fatal(r)
	}
}
