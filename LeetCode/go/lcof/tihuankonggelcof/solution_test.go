package tihuankonggelcof

import "testing"

func TestFunc(t *testing.T) {
	s := "We are happy."
	r := replaceSpace(s)
	if r != "We%20are%20happy." {
		t.Fatal(r)
	}
}
