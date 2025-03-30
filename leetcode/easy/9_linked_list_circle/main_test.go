package linkedlistcircle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	nodes1 := []ListNode{
		{Val: 3},
		{Val: 2},
		{Val: 0},
		{Val: -4},
	}
	nodes1[0].Next = &nodes1[1]
	nodes1[1].Next = &nodes1[2]
	nodes1[2].Next = &nodes1[3]
	nodes1[3].Next = &nodes1[1]
	t.Run("1Hash", func(t *testing.T) {
		res := hasCycleHash(&nodes1[0])
		assert.Equal(t, true, res)
	})
	// t.Run("1FastSlowPTR", func(t *testing.T) {
	// 	res := hasCycle(&nodes1[0])
	// 	assert.Equal(t, true, res)
	// })

	nodes2 := []ListNode{
		{Val: 1},
		{Val: 2},
	}
	nodes2[0].Next = &nodes2[1]
	nodes2[1].Next = &nodes2[0]
	t.Run("2Hash", func(t *testing.T) {
		res := hasCycleHash(&nodes2[0])
		assert.Equal(t, true, res)
	})

	nodes3 := []ListNode{
		{Val: 1},
	}
	t.Run("3Hash", func(t *testing.T) {
		res := hasCycleHash(&nodes3[0])
		assert.Equal(t, false, res)
	})
}
