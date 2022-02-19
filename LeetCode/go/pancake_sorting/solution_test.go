package pancakesorting

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	arr := []int{3, 2, 4, 1}
	r := pancakeSort(arr)
	fmt.Println(r)

	arr = []int{1, 2, 3}
	r = pancakeSort(arr)
	fmt.Println(r)
}
