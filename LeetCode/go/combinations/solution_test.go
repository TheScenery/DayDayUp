package combinations

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	n := 4
	k := 2
	r := combine(n, k)
	fmt.Println(r)

	n = 1
	k = 1
	r = combine(n, k)
	fmt.Println(r)

	n = 5
	k = 2
	r = combine(n, k)
	fmt.Println(r)
}

func TestFunc1(t *testing.T) {
	n := 4
	k := 2
	r := combine1(n, k)
	fmt.Println(r)

	n = 1
	k = 1
	r = combine1(n, k)
	fmt.Println(r)

	n = 5
	k = 2
	r = combine1(n, k)
	fmt.Println(r)
}
