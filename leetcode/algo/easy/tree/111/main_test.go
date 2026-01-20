package treemindepth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeMinDepth(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 3},
			{Val: 9},
			{Val: 20},
			{Val: 15},
			{Val: 2},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]
		tree[2].Left = &tree[3]
		tree[2].Right = &tree[4]

		res := minDepth(&tree[0])

		assert.Equal(t, 2, res)
	})

	t.Run("2", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
			{Val: 4},
			{Val: 5},
		}
		tree[0].Right = &tree[1]
		tree[1].Right = &tree[2]
		tree[2].Right = &tree[3]
		tree[3].Right = &tree[4]

		res := minDepth(&tree[0])

		assert.Equal(t, 5, res)
	})

	t.Run("3", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 1},
		}

		res := minDepth(&tree[0])

		assert.Equal(t, 1, res)
	})

	t.Run("4", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 3},
			{Val: 9},
			{Val: 20},
			{Val: 15},
			{Val: 2},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]
		tree[1].Left = &tree[3]
		tree[1].Right = &tree[4]

		res := minDepth(&tree[0])

		assert.Equal(t, 2, res)
	})

	t.Run("5", func(t *testing.T) {
		res := minDepth(nil)

		assert.Equal(t, 0, res)
	})

}
