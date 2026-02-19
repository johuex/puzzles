package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := checkHeight(root)
	return balanced
}

func checkHeight(node *TreeNode) (int, bool) {
	if node == nil { // достигли конца
		return 0, true
	}

	leftH, leftBal := checkHeight(node.Left)
	if !leftBal {
		return 0, false
	}

	rightH, rightBal := checkHeight(node.Right)
	if !rightBal {
		return 0, false
	}

	if abs(leftH-rightH) > 1 {
		return 0, false
	}

	return max(leftH, rightH) + 1, true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
