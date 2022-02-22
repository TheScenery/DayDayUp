package lettercasepermutation

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	s := "a1b2"
	r := letterCasePermutation(s)
	fmt.Println(r)
}
