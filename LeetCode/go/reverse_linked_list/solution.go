package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(parent, l *ListNode) *ListNode {
	if l == nil {
		return parent
	}
	lNext := l.Next
	l.Next = parent
	return reverse(l, lNext)
}

func reverseList(head *ListNode) *ListNode {
	return reverse(nil, head)
}
