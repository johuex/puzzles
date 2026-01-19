package deleteduplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeLinkedList(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		// 1,2,6,3,4,5,6
		nodes0 := []ListNode{
			{Val: 1},
			{Val: 1},
			{Val: 2},
		}
		nodes0[0].Next = &nodes0[1]
		nodes0[1].Next = &nodes0[2]

		res := deleteDuplicates(&nodes0[0])

		nodes1 := []ListNode{
			{Val: 1},
			{Val: 2},
		}
		nodes1[0].Next = &nodes1[1]

		expected := &nodes1[0]

		for range 2 {
			assert.Equal(t, res.Val, expected.Val)
			res = res.Next
			expected = expected.Next
		}
	})

	t.Run("2", func(t *testing.T) {
		// 1,1,2,3,3
		nodes0 := []ListNode{
			{Val: 1},
			{Val: 1},
			{Val: 2},
			{Val: 3},
			{Val: 3},
		}
		nodes0[0].Next = &nodes0[1]
		nodes0[1].Next = &nodes0[2]
		nodes0[2].Next = &nodes0[3]
		nodes0[3].Next = &nodes0[4]

		res := deleteDuplicates(&nodes0[0])

		nodes1 := []ListNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
		}
		nodes1[0].Next = &nodes1[1]
		nodes1[1].Next = &nodes1[2]

		expected := &nodes1[0]

		for range 3 {
			assert.Equal(t, res.Val, expected.Val)
			res = res.Next
			expected = expected.Next
		}
	})
}
