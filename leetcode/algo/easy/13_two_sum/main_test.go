package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{2, 7, 11, 15}, 9}, "output": []int{0, 1}},
	{"name": "2", "input": []any{[]int{3, 2, 4}, 6}, "output": []int{1, 2}},
	{"name": "3", "input": []any{[]int{3, 3}, 6}, "output": []int{0, 1}},
}

func TestTwoSum(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := twoSum(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
