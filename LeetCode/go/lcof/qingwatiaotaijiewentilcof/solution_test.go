package qingwatiaotaijiewentilcof

import "testing"

func TestFunc(t *testing.T) {
	r := numWays(7)
	if r != 21 {
		t.Fatal(r)
	}
}
