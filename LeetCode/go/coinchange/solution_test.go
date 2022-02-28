package coinchange

import "testing"

func TestFunc(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	r := coinChange(coins, amount)
	if r != 3 {
		t.Fatal(r)
	}

	coins = []int{2}
	amount = 3
	r = coinChange(coins, amount)
	if r != -1 {
		t.Fatal(r)
	}

	coins = []int{1}
	amount = 0
	r = coinChange(coins, amount)
	if r != 0 {
		t.Fatal(r)
	}
}
