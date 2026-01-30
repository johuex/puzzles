package jumpgame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name":   "1",
		"input":  []any{[]int{2, 3, 1, 1, 4}},
		"output": true,
	},
	{
		"name":   "2",
		"input":  []any{[]int{3, 2, 1, 0, 4}},
		"output": false,
	},
	{
		"name":   "3",
		"input":  []any{[]int{2, 5, 0, 0}},
		"output": true,
	},
	{
		"name":   "4",
		"input":  []any{[]int{1, 1, 2, 2, 0, 1, 1}},
		"output": true,
	},
	{
		"name":   "5",
		"input":  []any{[]int{0, 2, 3}},
		"output": false,
	},
	{
		"name":   "6",
		"input":  []any{[]int{1, 0, 2}},
		"output": false,
	},
	{
		"name":   "7",
		"input":  []any{[]int{0}},
		"output": true,
	},
	{
		"name":   "8",
		"input":  []any{[]int{0, 1}},
		"output": false,
	},
	{
		"name":   "9",
		"input":  []any{[]int{1, 0}},
		"output": true,
	},
	{
		"name":   "10",
		"input":  []any{[]int{1, 1}},
		"output": true,
	},
	{
		"name":   "11",
		"input":  []any{[]int{1, 0, 1, 0}},
		"output": false,
	},
	{
		"name":   "12",
		"input":  []any{[]int{1, 0, 0, 0, 3, 0, 0, 1}},
		"output": false,
	},
}

func TestJumpgame(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := canJump(testCase["input"].([]any)[0].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
