package middlelistnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MiddleListTestFunc(t *testing.T, testFunc func(*ListNode) *ListNode, name string) {
	nodes0 := []ListNode{
		{Val: 1},
		{Val: 2},
		{Val: 3},
		{Val: 4},
		{Val: 5},
		{Val: 6},
	}
	nodes0[0].Next = &nodes0[1]
	nodes0[1].Next = &nodes0[2]
	nodes0[2].Next = &nodes0[3]
	nodes0[3].Next = &nodes0[4]
	nodes0[4].Next = &nodes0[5]
	t.Run(name+": 0", func(t *testing.T) {
		res := testFunc(&nodes0[0])
		assert.Equal(t, &nodes0[3], res)
	})

	nodes1 := []ListNode{
		{Val: 1},
		{Val: 2},
		{Val: 3},
		{Val: 4},
		{Val: 5},
	}
	nodes1[0].Next = &nodes1[1]
	nodes1[1].Next = &nodes1[2]
	nodes1[2].Next = &nodes1[3]
	nodes1[3].Next = &nodes1[4]
	t.Run(name+": 1", func(t *testing.T) {
		res := testFunc(&nodes1[0])
		assert.Equal(t, &nodes1[2], res)
	})

	nodes2 := []ListNode{
		{Val: 1},
		{Val: 2},
	}
	nodes2[0].Next = &nodes2[1]
	t.Run(name+": 2", func(t *testing.T) {
		res := testFunc(&nodes2[0])
		assert.Equal(t, &nodes2[1], res)
	})

	nodes3 := []ListNode{
		{Val: 1},
	}
	t.Run(name+": 3", func(t *testing.T) {
		res := testFunc(&nodes3[0])
		assert.Equal(t, &nodes3[0], res)
	})
}

func TestMiddleList(t *testing.T) {
	MiddleListTestFunc(t, middleNode, "list method")
	MiddleListTestFunc(t, middleNodePointers, "two pointers method")
}
