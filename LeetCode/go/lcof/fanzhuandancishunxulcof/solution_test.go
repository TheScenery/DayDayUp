package fanzhuandancishunxulcof

import (
	"testing"
)

func TestFunc(t *testing.T) {
	s := "  hello world!  "
	r := reverseWords(s)
	if r != "world! hello" {
		t.Fatal(r)
	}
}
