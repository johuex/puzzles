package diameterofbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeMinDepth(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
			{Val: 4},
			{Val: 5},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]
		tree[1].Left = &tree[3]
		tree[1].Right = &tree[4]

		res := diameterOfBinaryTree(&tree[0])

		assert.Equal(t, 3, res)
	})

	t.Run("2", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 1},
			{Val: 2},
		}
		tree[0].Right = &tree[1]

		res := diameterOfBinaryTree(&tree[0])

		assert.Equal(t, 1, res)
	})

	t.Run("3", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 1},
		}

		res := diameterOfBinaryTree(&tree[0])

		assert.Equal(t, 0, res)
	})

	t.Run("4", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
			{Val: 4},
			{Val: 5},
			{Val: 6},
			{Val: 7},
			{Val: 8},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]
		tree[1].Left = &tree[3]
		tree[1].Right = &tree[4]
		tree[4].Left = &tree[5]
		tree[5].Right = &tree[6]
		tree[5].Right = &tree[7]

		res := diameterOfBinaryTree(&tree[0])

		assert.Equal(t, 5, res)
	})

	t.Run("5", func(t *testing.T) {
		tree := []TreeNode{
			{Val: 1},
			{Val: 2},
			{Val: 3},
			{Val: 4},
			{Val: 5},
			{Val: 6},
			{Val: 7},
			{Val: 8},
			{Val: 9},
			{Val: 10},
		}
		tree[0].Left = &tree[1]
		tree[0].Right = &tree[2]
		tree[1].Left = &tree[3]
		tree[1].Right = &tree[4]
		tree[3].Right = &tree[8]
		tree[8].Right = &tree[9]
		tree[4].Left = &tree[5]
		tree[5].Right = &tree[6]
		tree[5].Right = &tree[7]

		res := diameterOfBinaryTree(&tree[0])

		assert.Equal(t, 6, res)
	})

	t.Run("6", func(t *testing.T) {
		res := diameterOfBinaryTree(nil)

		assert.Equal(t, 0, res)
	})

}
