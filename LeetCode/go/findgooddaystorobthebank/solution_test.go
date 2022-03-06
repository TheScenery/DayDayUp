package findgooddaystorobthebank

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	security := []int{5, 3, 3, 3, 5, 6, 2}
	time := 2
	r := goodDaysToRobBank(security, time)
	fmt.Println(r)

	security = []int{1, 1, 1, 1, 1}
	time = 0
	r = goodDaysToRobBank(security, time)
	fmt.Println(r)

	security = []int{1, 2, 3, 4, 5, 6}
	time = 2
	r = goodDaysToRobBank(security, time)
	fmt.Println(r)

	security = []int{1}
	time = 5
	r = goodDaysToRobBank(security, time)
	fmt.Println(r)
}
