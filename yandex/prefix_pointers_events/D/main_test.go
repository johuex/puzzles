package d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			[]int{5, 1, 1, 1, 1},
		},
		"output": []int{1, 1, 2},
	},
	{
		"name": "2",
		"input": []any{
			[]int{1, 2, 3, 4},
		},
		"output": []int{1, 2, 4},
	},
	{
		"name": "3",
		"input": []any{
			[]int{1000, 399, 100, 100, 100, 100, 100, 100, 100, 100},
		},
		"output": []int{199, 1, 2},
	},
}

func TestFriends(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := friends(
				testCase["input"].([]any)[0].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
