package zuoxuanzhuanzifuchuanlcof

import "testing"

func TestFunc(t *testing.T) {
	s := "abcdefg"
	k := 2
	e := "cdefgab"
	r := reverseLeftWords(s, k)
	if r != e {
		t.Fatal(r)
	}

	s = "lrloseumgh"
	k = 6
	e = "umghlrlose"
	r = reverseLeftWords(s, k)
	if r != e {
		t.Fatal(r)
	}
}
