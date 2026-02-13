package j

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{
		[]int{0, 1, 1, 2, 3, 3},
		[][]int{
			{4, 1},
			{1, 4},
			{3, 6},
			{2, 6},
			{6, 5},
		},
	},
		"output": []int{0, 1, 1, 0, 0},
	},
	{"name": "2", "input": []any{
		[]int{0},
		[][]int{
			{1, 1},
		},
	},
		"output": []int{1},
	},
	{"name": "3", "input": []any{
		[]int{5, 4, 0, 3, 4, 7, 3},
		[][]int{
			{2, 7},
			{2, 7},
			{5, 1},
			{2, 4},
			{7, 2},
			{2, 6},
			{6, 2},
			{4, 2},
			{5, 1},
			{1, 7},
		},
	},
		"output": []int{0, 0, 1, 0, 0, 0, 0, 1, 1, 0},
	},
}

func TestParent(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := parent(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].([][]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
