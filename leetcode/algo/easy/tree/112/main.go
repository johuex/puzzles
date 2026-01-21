package treeavglevel

import "container/list"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueElem struct {
	Node *TreeNode
	Sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		// corner case -- empty root
		return false
	}

	queue := list.New()
	queue.PushBack(QueueElem{
		Node: root,
		Sum:  0,
	})

	for queue.Len() > 0 {
		// get elem from queue
		q := queue.Front()
		e := q.Value.(QueueElem)
		queue.Remove(q)
		node := e.Node
		currSum := e.Sum

		// we need to check target sum if it's leave
		if node.Left == nil && node.Right == nil {
			if currSum+node.Val == targetSum {
				return true
			}
			continue // it's leave, no childrens to add to queue
		}

		// push back childrens to queue
		if node.Left != nil {
			queue.PushBack(QueueElem{
				Node: node.Left,
				Sum:  currSum + node.Val,
			})
		}
		if node.Right != nil {
			queue.PushBack(QueueElem{
				Node: node.Right,
				Sum:  currSum + node.Val,
			})
		}
	}

	return false
}
