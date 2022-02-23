package populatingnextrightpointersineachnode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func setNext(n *Node, next *Node) {
	if n == nil {
		return
	}
	n.Next = next
	setNext(n.Left, n.Right)
	var rightNext *Node
	if n.Next != nil {
		rightNext = n.Next.Left
	}
	setNext(n.Right, rightNext)
}

func connect(root *Node) *Node {
	setNext(root, nil)
	return root
}
