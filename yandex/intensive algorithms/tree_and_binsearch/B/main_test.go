package b

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{
		5,
		[][]int{
			{1, 2},
			{1, 3},
			{2, 4},
			{2, 5},
		},
	},
		"output": 2,
	},
	{"name": "2", "input": []any{
		5,
		[][]int{
			{1, 3},
			{2, 1},
			{4, 5},
			{5, 3},
		},
	},
		"output": 4,
	},
}

func TestParent(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := logic(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].([][]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
