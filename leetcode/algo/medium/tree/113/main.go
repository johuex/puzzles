package pathsum2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	var dfs func(node *TreeNode, path []int, sum int)

	dfs = func(node *TreeNode, path []int, sum int) {
		if node == nil {
			return
		}

		// добавляем текущий узел
		path = append(path, node.Val)
		sum += node.Val

		// листок
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				tmp := make([]int, len(path))
				copy(tmp, path)
				result = append(result, tmp)
			}
			return
		}

		// идем вглубь в детей
		dfs(node.Left, path, sum)
		dfs(node.Right, path, sum)
	}

	dfs(root, []int{}, 0)
	return result
}
