package longestincseqlength

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name":   "1",
		"input":  []any{[]int{1, 3, 5, 4, 7}},
		"output": 3,
	},
	{
		"name":   "2",
		"input":  []any{[]int{2, 2, 2, 2, 2}},
		"output": 1,
	},
	{
		"name":   "3",
		"input":  []any{[]int{1}},
		"output": 1,
	},
}

func TestLongestLIS(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := findLengthOfLCIS(testCase["input"].([]any)[0].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
