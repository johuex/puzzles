package houserobber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name":   "1",
		"input":  []any{[]int{1, 2, 3, 1}},
		"output": 4,
	},
	{
		"name":   "2",
		"input":  []any{[]int{2, 7, 9, 3, 1}},
		"output": 12,
	},
	{
		"name":   "3",
		"input":  []any{[]int{40, 2, 4, 10}},
		"output": 50,
	},
	{
		"name":   "4",
		"input":  []any{[]int{9, 8, 0, 8, 9}},
		"output": 18,
	},
	{
		"name":   "5",
		"input":  []any{[]int{9, 8, 8, 9}},
		"output": 18,
	},
}

func TestRob(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := rob(testCase["input"].([]any)[0].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
