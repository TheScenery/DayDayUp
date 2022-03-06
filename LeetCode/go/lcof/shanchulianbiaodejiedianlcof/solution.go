package shanchulianbiaodejiedianlcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	pre := head
	cur := head.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			return head
		}
		pre, cur = cur, cur.Next
	}
	return head
}
