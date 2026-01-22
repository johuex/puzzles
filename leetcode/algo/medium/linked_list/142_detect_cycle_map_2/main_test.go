package linkedlistcycle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T) {
	nodes1 := []ListNode{
		{Val: 3},
		{Val: 2},
		{Val: 0},
		{Val: 4},
	}
	nodes1[0].Next = &nodes1[1]
	nodes1[1].Next = &nodes1[2]
	nodes1[2].Next = &nodes1[3]
	nodes1[3].Next = &nodes1[1]
	t.Run("1Hash", func(t *testing.T) {
		res := detectCycleMap(&nodes1[0])
		assert.Equal(t, &nodes1[1], res)
	})
	t.Run("1FastSlowPTR", func(t *testing.T) {
		res := detectCyclePointers(&nodes1[0])
		assert.Equal(t, &nodes1[1], res)
	})

	nodes2 := []ListNode{
		{Val: 1},
		{Val: 2},
	}
	nodes2[0].Next = &nodes2[1]
	nodes2[1].Next = &nodes2[0]
	t.Run("2Hash", func(t *testing.T) {
		res := detectCycleMap(&nodes2[0])
		assert.Equal(t, &nodes2[0], res)
	})
	t.Run("2FastSlowPTR", func(t *testing.T) {
		res := detectCyclePointers(&nodes2[0])
		assert.Equal(t, &nodes2[0], res)
	})

	nodes3 := []ListNode{
		{Val: 1},
	}
	t.Run("3Hash", func(t *testing.T) {
		res := detectCycleMap(&nodes3[0])
		assert.Nil(t, res)
	})
	t.Run("3FastSlowPTR", func(t *testing.T) {
		res := detectCyclePointers(&nodes3[0])
		assert.Nil(t, res)
	})

	t.Run("4Hash", func(t *testing.T) {
		res := detectCycleMap(nil)
		assert.Nil(t, res)
	})
	t.Run("4FastSlowPTR", func(t *testing.T) {
		res := detectCyclePointers(nil)
		assert.Nil(t, res)
	})
}
