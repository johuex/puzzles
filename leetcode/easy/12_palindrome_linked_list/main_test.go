package palindrominkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func palindromeListTestFunc(t *testing.T, testFunc func(*ListNode) bool, name string) {
	nodes0 := []ListNode{
		{Val: 1},
		{Val: 2},
		{Val: 3},
		{Val: 2},
		{Val: 1},
	}
	nodes0[0].Next = &nodes0[1]
	nodes0[1].Next = &nodes0[2]
	nodes0[2].Next = &nodes0[3]
	nodes0[3].Next = &nodes0[4]
	t.Run(name+": 0", func(t *testing.T) {
		res := testFunc(&nodes0[0])
		assert.Equal(t, true, res)
	})

	nodes1 := []ListNode{
		{Val: 1},
		{Val: 2},
		{Val: 2},
		{Val: 1},
	}
	nodes1[0].Next = &nodes1[1]
	nodes1[1].Next = &nodes1[2]
	nodes1[2].Next = &nodes1[3]
	t.Run(name+": 1", func(t *testing.T) {
		res := testFunc(&nodes1[0])
		assert.Equal(t, true, res)
	})

	nodes2 := []ListNode{
		{Val: 1},
		{Val: 2},
	}
	nodes2[0].Next = &nodes2[1]
	t.Run(name+": 2", func(t *testing.T) {
		res := testFunc(&nodes2[0])
		assert.Equal(t, false, res)
	})

	nodes3 := []ListNode{
		{Val: 1},
	}
	t.Run(name+": 3", func(t *testing.T) {
		res := testFunc(&nodes3[0])
		assert.Equal(t, true, res)
	})

	nodes1 = []ListNode{
		{Val: 1},
		{Val: 2},
		{Val: 6},
		{Val: 1},
	}
	nodes1[0].Next = &nodes1[1]
	nodes1[1].Next = &nodes1[2]
	nodes1[2].Next = &nodes1[3]
	t.Run(name+": 4", func(t *testing.T) {
		res := testFunc(&nodes1[0])
		assert.Equal(t, false, res)
	})
}

func TestPalindromeList(t *testing.T) {
	//palindromeListTestFunc(t, isPalindrome, "with list")
	palindromeListTestFunc(t, isPalindromePointers, "with pointers")
}
