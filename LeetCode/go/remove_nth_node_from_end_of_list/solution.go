package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	nthNode := head
	result := head
	for n >= 0 && nthNode != nil {
		nthNode = nthNode.Next
		n--
	}
	if n >= 0 {
		result = head.Next
	}

	for nthNode != nil {
		current = current.Next
		nthNode = nthNode.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
	}
	return result
}
