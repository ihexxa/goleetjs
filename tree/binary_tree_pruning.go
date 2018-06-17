package tree


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	found := foundOne(root)
	if !found {
		root = nil
	}

	return root
}

// test not added
func foundOne(root *TreeNode) bool {
	lFound := false
	rFound := false

	if root.Left != nil {
		lFound = foundOne(root.Left)
		if !lFound {
			root.Left = nil
		}
	}
	if root.Right != nil {
		rFound = foundOne(root.Right)
		if !rFound {
			root.Right = nil
		}
	}

	return root.Val == 1 || lFound || rFound
}
