package addtwonumbers

import (
	"fmt"
	"testing"
)

func (l *ListNode) FromArray(a []int) {
	currentL := l
	length := len(a)
	for i, v := range a {
		currentL.Val = v
		if i != length-1 {
			currentL.Next = &ListNode{}
			currentL = currentL.Next
		}
	}
}

func (l *ListNode) ToArray() []int {
	result := []int{}
	currentL := l
	for {
		result = append(result, currentL.Val)
		if currentL.Next == nil {
			break
		}
		currentL = currentL.Next
	}
	return result
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{}
	l1.FromArray([]int{2, 4, 3})

	l2 := &ListNode{}
	l2.FromArray([]int{5, 6, 4})

	result := addTwoNumbers(l1, l2)
	fmt.Println(result.ToArray())
}
