package treeavglevel

import "container/list"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueElem struct {
	Node  *TreeNode
	Level int
}

func averageOfLevels(root *TreeNode) []float64 {
	queue := list.New()
	levelMap := make(map[int][]int)
	currLevel := 0
	queue.PushBack(QueueElem{
		Node:  root,
		Level: currLevel,
	})

	for queue.Len() > 0 {
		// get elem from queue
		q := queue.Front()
		e := q.Value.(QueueElem)
		currLevel = e.Level
		queue.Remove(q) // TODO: или в конце удалять???
		node := e.Node

		// append elem value to level
		levelArr, ok := levelMap[currLevel]
		if !ok {
			levelArr = make([]int, 0)
		}
		levelArr = append(levelArr, node.Val)
		levelMap[currLevel] = levelArr

		// push back childrens to queue
		if node.Left != nil {
			queue.PushBack(QueueElem{
				Node:  node.Left,
				Level: e.Level + 1,
			})
		}
		if node.Right != nil {
			queue.PushBack(QueueElem{
				Node:  node.Right,
				Level: e.Level + 1,
			})
		}
	}

	res := make([]float64, len(levelMap))

	// calculate AVG for all levels
	for key, val := range levelMap {
		sum := 0
		for _, v := range val {
			sum += v
		}
		avg := float64(sum) / float64(len(val))
		res[key] = avg
	}

	return res
}
