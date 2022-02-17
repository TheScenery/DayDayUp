package permutationinstring

import "testing"

func TestFunc(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	r := checkInclusion(s1, s2)
	if r != true {
		t.Fatal(r)
	}

	s1 = "ab"
	s2 = "eidboaoo"
	r = checkInclusion(s1, s2)
	if r != false {
		t.Fatal(r)
	}

	s1 = "adc"
	s2 = "dcda"
	r = checkInclusion(s1, s2)
	if r != true {
		t.Fatal(r)
	}
}

func TestFunc1(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	r := checkInclusion1(s1, s2)
	if r != true {
		t.Fatal(r)
	}

	s1 = "ab"
	s2 = "eidboaoo"
	r = checkInclusion1(s1, s2)
	if r != false {
		t.Fatal(r)
	}

	s1 = "adc"
	s2 = "dcda"
	r = checkInclusion1(s1, s2)
	if r != true {
		t.Fatal(r)
	}
}
