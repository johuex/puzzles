package g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			[]int{6, 7, 8, 8, 7, 7},
			[]int{10, 6, 3, 1, 1, 4},
		},
		"output": []int{7, 14},
	},
	{
		"name": "2",
		"input": []any{
			[]int{7, 5, 7, 9, 8},
			[]int{10, 8, 7, 8, 5},
		},
		"output": []int{7, 37},
	},
	{
		"name": "3",
		"input": []any{
			[]int{8, 5, 10, 9, 7},
			[]int{2, 5, 4, 8, 4},
		},
		"output": []int{9, 34},
	},
}

func TestCompensation(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := compensation(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
