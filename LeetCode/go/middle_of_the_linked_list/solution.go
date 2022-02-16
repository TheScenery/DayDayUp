package middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	listCount := 0
	currentNode := head
	for currentNode != nil {
		listCount++
		currentNode = currentNode.Next
	}
	mid := listCount / 2
	currentNode = head
	for mid >= 0 {
		currentNode = currentNode.Next
		mid--
	}
	return currentNode
}

func middleNode1(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return slow
}
