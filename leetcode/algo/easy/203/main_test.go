package removeelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeLinkedList(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		// 1,2,6,3,4,5,6
		nodes0 := []ListNode{
			{Val: 6},
			{Val: 1},
			{Val: 2},
			{Val: 6},
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
		nodes0[5].Next = &nodes0[6]
		nodes0[6].Next = &nodes0[7]

		res := removeElements(&nodes0[0], 6)

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

		expected := &nodes1[0]

		for range 5 {
			assert.Equal(t, res.Val, expected.Val)
			res = res.Next
			expected = expected.Next
		}
	})
}
