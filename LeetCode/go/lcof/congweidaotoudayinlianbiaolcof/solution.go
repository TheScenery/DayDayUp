package congweidaotoudayinlianbiaolcof

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	result := make([]int, 0)
	var visit func(*ListNode)
	visit = func(p *ListNode) {
		if p == nil {
			return
		}
		visit(p.Next)
		result = append(result, p.Val)
	}
	visit(head)
	return result
}
