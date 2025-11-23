package e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			[]int{1, 2, 3, 4},
			[][]int{
				{1, 4},
				{3, 4},
			},
			2,
		},
		"output": 13,
	},
	{
		"name": "2",
		"input": []any{
			[]int{1, 2, 0, 0},
			[][]int{
				{1, 4},
				{3, 4},
			},
			5,
		},
		"output": 0,
	},
}

func TestFriends(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := repair(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].([][]int),
				testCase["input"].([]any)[2].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
