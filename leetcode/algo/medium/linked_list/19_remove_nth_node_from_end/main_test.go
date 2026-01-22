package removenelem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// DISCLAIMER:
// don't want to write huge assert part, so use it to laucnh and check manually

func TestRemoveNth(t *testing.T) {
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
	t.Run("1Hash", func(t *testing.T) {
		res := removeNthFromEnd(&nodes1[0], 2)
		assert.Equal(t, &nodes1[1], res)
	})

}
