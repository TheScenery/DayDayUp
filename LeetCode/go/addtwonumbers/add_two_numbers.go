package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	header := &ListNode{}
	current := header
	currentL1 := l1
	currentL2 := l2
	bit := 0
	for {
		currentValue := currentL1.Val + currentL2.Val + bit
		if currentValue >= 10 {
			bit = 1
			current.Val = currentValue - 10
		} else {
			bit = 0
			current.Val = currentValue
		}
		if bit == 0 && currentL1.Next == nil && currentL2.Next == nil {
			break
		}

		current.Next = &ListNode{}
		current = current.Next
		if currentL1.Next != nil {
			currentL1 = currentL1.Next
		} else {
			currentL1 = &ListNode{}
		}

		if currentL2.Next != nil {
			currentL2 = currentL2.Next
		} else {
			currentL2 = &ListNode{}
		}
	}
	return header
}
