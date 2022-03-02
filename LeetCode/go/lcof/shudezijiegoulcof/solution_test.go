package shudezijiegoulcof

import (
	"testing"
)

func TestFunc(t *testing.T) {
	A := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	B := &TreeNode{
		Val: 3,
	}
	r := isSubStructure(A, B)
	if r != true {
		t.Fatal()
	}
}
