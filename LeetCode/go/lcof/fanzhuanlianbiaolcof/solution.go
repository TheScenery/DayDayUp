package fanzhuanlianbiaolcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	newHead := head
	var visit func(*ListNode)
	visit = func(p *ListNode) {
		if p == nil || p.Next == nil {
			newHead = p
			return
		}
		visit(p.Next)
		p.Next.Next = p
		if p == head {
			p.Next = nil
		}

	}
	visit(head)
	return newHead
}
