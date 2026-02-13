package turismhigh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			[][]int{
				{2, 1},
				{4, 5},
				{7, 4},
				{8, 2},
				{9, 6},
				{11, 3},
				{15, 3},
			},
			[][]int{{2, 6}},
		},
		"output": []int{4},
	},
	{
		"name": "2",
		"input": []any{
			[][]int{
				{1, 1},
				{3, 2},
				{5, 6},
				{7, 2},
				{10, 4},
				{11, 1},
			},
			[][]int{
				{5, 6},
				{1, 4},
				{4, 2},
			},
		},
		"output": []int{0, 5, 4},
	},
}

func TestNearestNumber(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := turism(testCase["input"].([]any)[0].([][]int), testCase["input"].([]any)[1].([][]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
