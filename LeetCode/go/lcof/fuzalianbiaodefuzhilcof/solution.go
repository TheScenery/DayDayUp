package fuzalianbiaodefuzhilcof

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyNode(n *Node, copiedNodes map[*Node]*Node) *Node {
	if n == nil {
		return n
	}
	if copiedNodes[n] != nil {
		return copiedNodes[n]
	}
	newNode := Node{
		Val: n.Val,
	}
	copiedNodes[n] = &newNode
	newNode.Next = copyNode(n.Next, copiedNodes)
	newNode.Random = copyNode(n.Random, copiedNodes)
	return &newNode
}

func copyRandomList(head *Node) *Node {
	copied := make(map[*Node]*Node)
	return copyNode(head, copied)
}
