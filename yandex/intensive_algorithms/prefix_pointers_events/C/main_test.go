package c

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			2,
			[]int{3},
			[][]int{
				{1, 2},
				{1, 1},
				{3, 0},
				{3, 1},
				{3, 2},
			},
		},
		"output": []int{
			0,
			1,
			2,
		},
	},
	{
		"name": "2",
		"input": []any{
			2,
			[]int{1, 2},
			[][]int{
				{3, 0},
				{3, 1},
				{2},
				{3, 0},
				{1, 3},
				{3, 0},
				{3, 1},
			},
		},
		"output": []int{
			0,
			0,
			0,
			0,
			1,
		},
	},
}

func TestRiver(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := candidatesCheck(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].([]int),
				testCase["input"].([]any)[2].([][]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
