package fibonaccinumber

import "testing"

func TestFunc(t *testing.T) {
	r := fib(2)
	if r != 1 {
		t.Fatal(r)
	}

	r = fib(3)
	if r != 2 {
		t.Fatal(r)
	}

	r = fib(4)
	if r != 3 {
		t.Fatal(r)
	}
}
