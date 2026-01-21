package treeavglevel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeMinDepth(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 5},
			{Val: 4},
			{Val: 8},
			{Val: 11},
			{Val: 7},
			{Val: 2},
			{Val: 13},
			{Val: 4},
			{Val: 1},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]
		tree[1].Left = &tree[3]
		tree[3].Left = &tree[4]
		tree[3].Right = &tree[5]
		tree[2].Left = &tree[6]
		tree[2].Right = &tree[7]
		tree[7].Right = &tree[8]

		res := hasPathSum(&tree[0], 22)

		assert.Equal(t, true, res)
	})

	t.Run("2", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]

		res := hasPathSum(&tree[0], 5)

		assert.Equal(t, false, res)
	})

	t.Run("3", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]

		res := hasPathSum(nil, 0)

		assert.Equal(t, false, res)
	})
}
