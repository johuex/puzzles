package searchinrotated

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, "output": 4},
	{"name": "2", "input": []any{[]int{4, 5, 6, 7, 0, 1, 2}, 3}, "output": -1},
	{"name": "3", "input": []any{[]int{1}, 0}, "output": -1},
	{"name": "4", "input": []any{[]int{}, 0}, "output": -1},
	{"name": "5", "input": []any{[]int{4, 5, 6, 7, 0, 1, 2}, 0}, "output": 4},
	{"name": "6", "input": []any{[]int{1, 2, 4, 5, 6, 7, 0}, 0}, "output": 6},
	{"name": "7", "input": []any{[]int{5, 6, 7, 0, 1, 2, 4}, 0}, "output": 3},
	{"name": "8", "input": []any{[]int{7, 0, 1, 2, 4, 5, 6}, 0}, "output": 1},
	{"name": "9", "input": []any{[]int{1, 3}, 2}, "output": -1},
	{"name": "10", "input": []any{[]int{1, 3, 5}, 3}, "output": 1},
}

func TestSearchTarget(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := search(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
