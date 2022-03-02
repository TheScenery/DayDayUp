package duichengdeerchashulcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSame(A, B *TreeNode) bool {
	if A == B {
		return true
	}
	if A == nil || B == nil {
		return false
	}
	return A.Val == B.Val && isSame(A.Left, B.Left) && isSame(A.Right, B.Right)
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  mirrorTree(root.Right),
		Right: mirrorTree(root.Left),
	}
}

func isSymmetric(root *TreeNode) bool {
	return isSame(root, mirrorTree(root))
}

func recur(L, R *TreeNode) bool {
	if L == R {
		return true
	}
	if L == nil || R == nil || L.Val != R.Val {
		return false
	}
	return recur(L.Left, R.Right) && recur(R.Left, L.Right)
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}
