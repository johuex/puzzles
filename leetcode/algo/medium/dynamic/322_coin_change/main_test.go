package coinchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name":   "1",
		"input":  []any{[]int{1, 2, 5}, 11},
		"output": 3,
	},
	{
		"name":   "2",
		"input":  []any{[]int{2}, 3},
		"output": -1,
	},
	{
		"name":   "3",
		"input":  []any{[]int{1}, 0},
		"output": 0,
	},
	{
		"name":   "4",
		"input":  []any{[]int{186, 419, 83, 408}, 6249},
		"output": 20,
	},
	//
	{
		"name":   "5",
		"input":  []any{[]int{1, 2, 5}, 8839},
		"output": 1769,
	},
	{
		"name":   "6",
		"input":  []any{[]int{3, 7, 405, 436}, 6249},
		"output": 20,
	},
	{
		"name":   "7",
		"input":  []any{[]int{176, 6, 366, 357, 484, 226, 1, 104, 160, 331}, 5557},
		"output": 13,
	},
	{
		"name":   "8",
		"input":  []any{[]int{19, 28, 176, 112, 30, 260, 491, 128, 70, 137, 253}, 8539},
		"output": 21,
	},
	{
		"name":   "9",
		"input":  []any{[]int{370, 417, 408, 156, 143, 434, 168, 83, 177, 280, 117}, 9953},
		"output": 24,
	},
	{
		"name":   "10",
		"input":  []any{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24}, 9999},
		"output": -1,
	},
	{
		"name":   "11",
		"input":  []any{[]int{2, 4, 6}, 9999},
		"output": -1,
	},
}

func TestCoinChange(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := coinChange(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
