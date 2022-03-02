package shudezijiegoulcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isLike(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	return A.Val == B.Val && isLike(A.Left, B.Left) && isLike(A.Right, B.Right)
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if A.Val == B.Val {
		r := isLike(A, B)
		if r {
			return r
		}
	}
	r := isSubStructure(A.Left, B)
	if r {
		return true
	}
	return isSubStructure(A.Right, B)
}
