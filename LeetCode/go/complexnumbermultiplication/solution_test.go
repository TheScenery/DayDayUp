package complexnumbermultiplication

import "testing"

func TestFunc(t *testing.T) {
	num1 := "1+1i"
	num2 := "1+1i"
	r := complexNumberMultiply(num1, num2)
	if r != "0+2i" {
		t.Fatal(r)
	}

	num1 = "1+-1i"
	num2 = "1+-1i"
	r = complexNumberMultiply(num1, num2)
	if r != "0+-2i" {
		t.Fatal(r)
	}
}
