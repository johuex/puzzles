package reorderlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		list := []*ListNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
			{Val: 4},
			{Val: 5},
		}

		list[0].Next = list[1]
		list[1].Next = list[2]
		list[2].Next = list[3]
		list[3].Next = list[4]

		reorderList(list[0])

		ans := []int{1, 5, 2, 4, 3}
		node := list[0]
		for i := range len(ans) {
			assert.Equal(t, ans[i], node.Val)
			node = node.Next
		}
	})

	t.Run("2", func(t *testing.T) {
		list := []*ListNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
			{Val: 4},
		}

		list[0].Next = list[1]
		list[1].Next = list[2]
		list[2].Next = list[3]

		reorderList(list[0])

		ans := []int{1, 4, 2, 3}
		node := list[0]
		for i := range len(ans) {
			assert.Equal(t, ans[i], node.Val)
			node = node.Next
		}
	})
}
