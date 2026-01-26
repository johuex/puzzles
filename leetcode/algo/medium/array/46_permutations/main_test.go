package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{1, 2, 3}}, "output": [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	{"name": "2", "input": []any{[]int{0, 1}}, "output": [][]int{{0, 1}, {1, 0}}},
	{"name": "3", "input": []any{[]int{1}}, "output": [][]int{{1}}},
}

func TestLetterCasePermutation(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := permute(testCase["input"].([]any)[0].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
