package erchashuzhongheweimouyizhidelujinglcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	path := []int{root.Val}
	sum := root.Val

	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		if tn.Left == nil && tn.Right == nil {
			if sum == target {
				r := make([]int, len(path))
				copy(r, path)
				result = append(result, r)
				return
			}
		}

		choice := []*TreeNode{tn.Left, tn.Right}
		for _, c := range choice {
			if c != nil {
				sum += c.Val
				path = append(path, c.Val)
				dfs(c)

				sum -= c.Val
				path = path[0 : len(path)-1]
			}
		}
	}
	dfs(root)
	return result
}
