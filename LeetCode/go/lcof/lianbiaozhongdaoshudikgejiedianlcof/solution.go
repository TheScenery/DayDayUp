package lianbiaozhongdaoshudikgejiedianlcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	cur := head
	var nextK *ListNode
	for i := 1; i < k; i++ {
		if cur == nil {
			return nil
		}
		cur = cur.Next
	}
	nextK = cur
	cur = head
	for nextK != nil && nextK.Next != nil {
		nextK = nextK.Next
		cur = cur.Next
	}
	return cur
}
