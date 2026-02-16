package gistogram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{7, 2, 1, 4, 5, 1, 3}}, "output": 8},
	{"name": "2", "input": []any{[]int{7, 2, 1, 5, 4, 1, 3}}, "output": 8},
	{"name": "3", "input": []any{[]int{1, 5, 4, 5, 1, 3, 3}}, "output": 12},
	{"name": "4", "input": []any{[]int{2, 2, 2, 2}}, "output": 8},
	{"name": "5", "input": []any{[]int{3, 3, 3, 1, 3, 3}}, "output": 9},
	{"name": "6", "input": []any{[]int{3, 3, 3, 4, 3, 3}}, "output": 18},
	{"name": "7", "input": []any{[]int{3, 3, 3, 4, 3, 2}}, "output": 15},
	{"name": "8", "input": []any{[]int{}}, "output": 0},
	{"name": "9", "input": []any{[]int{2, 1, 4, 5, 1, 3, 3}}, "output": 8},
}

func TestGistogram(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := gistogram(testCase["input"].([]any)[0].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
