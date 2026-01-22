package longestsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name":   "1",
		"input":  []any{[]int{100, 4, 200, 1, 3, 2}},
		"output": 4,
	},
	{
		"name":   "2",
		"input":  []any{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
		"output": 9,
	},
	{
		"name":   "3",
		"input":  []any{[]int{1, 2, 0, 1}},
		"output": 3,
	},
	{
		"name":   "4",
		"input":  []any{[]int{0, 0}},
		"output": 1,
	},
	{
		"name":   "5",
		"input":  []any{[]int{}},
		"output": 0,
	},
	{
		"name":   "6",
		"input":  []any{[]int{1, 1, 2, 2, 2}},
		"output": 2,
	},
}

func TestFindDuplicate(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := longestConsecutive(testCase["input"].([]any)[0].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
