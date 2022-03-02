package congshangdaoxiadayinerchashulcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	currentLevelNodes := []*TreeNode{root}
	for len(currentLevelNodes) != 0 {
		node := currentLevelNodes[0]
		result = append(result, node.Val)
		currentLevelNodes = currentLevelNodes[1:]
		if node.Left != nil {
			currentLevelNodes = append(currentLevelNodes, node.Left)
		}
		if node.Right != nil {
			currentLevelNodes = append(currentLevelNodes, node.Right)
		}
	}
	return result
}
