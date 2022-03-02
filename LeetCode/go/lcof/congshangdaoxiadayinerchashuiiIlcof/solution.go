package congshangdaoxiadayinerchashuiiilcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	currentLevelNodes := []*TreeNode{root}
	i := 0
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
		if i%2 != 0 {
			reverse(r)
		}
		result = append(result, r)
		currentLevelNodes = nextLevelNodes
		i++
	}
	return result
}
