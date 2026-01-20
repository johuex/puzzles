package treeavglevel

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
			{Val: 7},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]
		tree[2].Left = &tree[3]
		tree[2].Right = &tree[4]

		res := averageOfLevels(&tree[0])

		assert.Equal(t, []float64{3.00000, 14.50000, 11.00000}, res)
	})

	t.Run("2", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 3},
			{Val: 9},
			{Val: 20},
			{Val: 15},
			{Val: 7},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]
		tree[1].Left = &tree[3]
		tree[1].Right = &tree[4]

		res := averageOfLevels(&tree[0])

		assert.Equal(t, []float64{3.00000, 14.50000, 11.00000}, res)
	})
}
