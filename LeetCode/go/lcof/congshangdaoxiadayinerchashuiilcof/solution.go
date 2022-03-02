package congshangdaoxiadayinerchashuiilcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	currentLevelNodes := []*TreeNode{root}
	for len(currentLevelNodes) != 0 {
		nextLevelNodes := make([]*TreeNode, 0)
		r := make([]int, 0)
		for _, node := range currentLevelNodes {
			r = append(r, node.Val)
			if node.Left != nil {
				nextLevelNodes = append(nextLevelNodes, node.Left)
			}
			if node.Right != nil {
				nextLevelNodes = append(nextLevelNodes, node.Right)
			}
		}
		result = append(result, r)
		currentLevelNodes = nextLevelNodes
	}
	return result
}
