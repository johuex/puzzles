package squaresorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{-4, -1, 0, 3, 10}}, "output": []int{0, 1, 9, 16, 100}},
	{"name": "2", "input": []any{[]int{-7, -3, 2, 3, 11}}, "output": []int{4, 9, 9, 49, 121}},
	{"name": "3", "input": []any{[]int{-5, -3, -2, -1}}, "output": []int{1, 4, 9, 25}},
	{"name": "4", "input": []any{[]int{1, 2, 3, 5}}, "output": []int{1, 4, 9, 25}},
}

func TestSortedSquares(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := sortedSquares(
				testCase["input"].([]any)[0].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
