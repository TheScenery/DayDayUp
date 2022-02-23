package reverseonlyletters

import "testing"

func TestFunc(t *testing.T) {
	s := "ab-cd"
	r := reverseOnlyLetters(s)
	if r != "dc-ba" {
		t.Fatal(r)
	}

	s = "a-bC-dEf-ghIj"
	r = reverseOnlyLetters(s)
	if r != "j-Ih-gfE-dCba" {
		t.Fatal(r)
	}

	s = "Test1ng-Leet=code-Q!"
	r = reverseOnlyLetters(s)
	if r != "Qedo1ct-eeLg=ntse-T!" {
		t.Fatal(r)
	}
}
