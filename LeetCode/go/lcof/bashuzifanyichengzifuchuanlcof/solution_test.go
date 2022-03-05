package bashuzifanyichengzifuchuanlcof

import "testing"

func TestFunc(t *testing.T) {
	num := 12258
	r := translateNum(num)
	if r != 5 {
		t.Fatal(r)
	}

	num = 25
	r = translateNum(num)
	if r != 2 {
		t.Fatal(r)
	}

	num = 58
	r = translateNum(num)
	if r != 1 {
		t.Fatal(r)
	}

	num = 506
	r = translateNum(num)
	if r != 1 {
		t.Fatal(r)
	}
}
