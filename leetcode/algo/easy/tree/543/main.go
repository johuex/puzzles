package diameterofbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxPath int // max tree diameter
	treeSearch(root, &maxPath)
	return maxPath
}

func treeSearch(root *TreeNode, maxPath *int) int {
	if root == nil {
		// corner start case, tree is empty
		return 0
	}

	if root.Left == nil && root.Right == nil {
		// it's leave, return path as 1
		return 1
	}
	var leftPath, rightPath int
	if root.Left != nil {
		leftPath = treeSearch(root.Left, maxPath)
	}

	if root.Right != nil {
		rightPath = treeSearch(root.Right, maxPath)
	}

	// check max path in node
	if leftPath+rightPath > *maxPath {
		*maxPath = leftPath + rightPath
	}

	return max(leftPath, rightPath) + 1
}
