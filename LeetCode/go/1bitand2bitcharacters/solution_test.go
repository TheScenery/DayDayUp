package onebitandtwobitcharacter

import "testing"

func TestFunc(t *testing.T) {
	bits := []int{1, 0, 0}
	r := isOneBitCharacter(bits)
	if r != true {
		t.Fatal(r)
	}

	bits = []int{1, 1, 1, 0}
	r = isOneBitCharacter(bits)
	if r != false {
		t.Fatal(r)
	}
}
