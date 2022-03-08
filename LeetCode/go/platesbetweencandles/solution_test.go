package platesbetweencandles

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	s := "**|**|***|"
	queries := [][]int{{2, 5}, {5, 9}}
	r := platesBetweenCandles(s, queries)
	fmt.Println(r)

	s = "***|**|*****|**||**|*"
	queries = [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}
	r = platesBetweenCandles(s, queries)
	fmt.Println(r)
}
