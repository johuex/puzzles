package maxsubarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name":   "1",
		"input":  []any{[]int{-2, -1, 4, -5, 7, -5}},
		"output": 7,
	},
	{
		"name":   "2",
		"input":  []any{[]int{-2, -1, 4, 2, -5, 6, 2, 3, -15}},
		"output": 12,
	},
	{
		"name":   "3",
		"input":  []any{[]int{1}},
		"output": 1,
	},
	{
		"name":   "4",
		"input":  []any{[]int{5, 4, -1, 7, 8}},
		"output": 23,
	},
}

func TestMaxSubarray(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := maxSubArray(testCase["input"].([]any)[0].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
