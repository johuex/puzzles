package mergelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeLinkedList(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nodes0 := []ListNode{
			{Val: 1},
			{Val: 3},
			{Val: 6},
		}
		nodes0[0].Next = &nodes0[1]
		nodes0[1].Next = &nodes0[2]

		nodes1 := []ListNode{
			{Val: 1},
			{Val: 2},
			{Val: 5},
		}
		nodes1[0].Next = &nodes1[1]
		nodes1[1].Next = &nodes1[2]

		res := mergeTwoLists(&nodes0[0], &nodes1[0])

		nodes2 := []ListNode{
			{Val: 1},
			{Val: 1},
			{Val: 2},
			{Val: 3},
			{Val: 5},
			{Val: 6},
		}
		nodes2[0].Next = &nodes2[1]
		nodes2[1].Next = &nodes2[2]
		nodes2[2].Next = &nodes2[3]
		nodes2[3].Next = &nodes2[4]
		nodes2[4].Next = &nodes2[5]

		expected := &nodes2[0]

		for range 6 {
			assert.Equal(t, res.Val, expected.Val)
			res = res.Next
			expected = expected.Next
		}
	})

	t.Run("2", func(t *testing.T) {
		res := mergeTwoLists(nil, nil)
		assert.Nil(t, res)
	})

	t.Run("3", func(t *testing.T) {
		nodes0 := []ListNode{
			{Val: 1},
			{Val: 3},
			{Val: 6},
		}
		nodes0[0].Next = &nodes0[1]
		nodes0[1].Next = &nodes0[2]

		res := mergeTwoLists(&nodes0[0], nil)

		expected := &nodes0[0]

		for range 3 {
			assert.Equal(t, res.Val, expected.Val)
			res = res.Next
			expected = expected.Next
		}
	})

	t.Run("4", func(t *testing.T) {
		nodes0 := []ListNode{
			{Val: 1},
			{Val: 3},
			{Val: 6},
		}
		nodes0[0].Next = &nodes0[1]
		nodes0[1].Next = &nodes0[2]

		res := mergeTwoLists(nil, &nodes0[0])

		expected := &nodes0[0]

		for range 3 {
			assert.Equal(t, res.Val, expected.Val)
			res = res.Next
			expected = expected.Next
		}
	})
}
