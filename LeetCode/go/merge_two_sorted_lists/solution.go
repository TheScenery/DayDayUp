package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	combinedNode := list1
	if list2.Val < list1.Val {
		combinedNode = list2
		combinedNode.Next = mergeTwoLists(list2.Next, list1)
	} else {
		combinedNode.Next = mergeTwoLists(list1.Next, list2)
	}

	return combinedNode
}
