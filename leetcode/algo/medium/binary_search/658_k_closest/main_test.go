package kclosest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{1, 2, 3, 4, 5}, 4, 3}, "output": []int{1, 2, 3, 4}},
	{"name": "2", "input": []any{[]int{1, 1, 2, 3, 4, 5}, 4, -1}, "output": []int{1, 1, 2, 3}},
	{"name": "3", "input": []any{[]int{1, 1, 2, 3, 4, 5}, 4, 6}, "output": []int{2, 3, 4, 5}},
	{"name": "4", "input": []any{[]int{1, 1, 1, 10, 10, 10}, 1, 9}, "output": []int{10}},
}

func TestFindClosestElements(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := findClosestElements(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].(int),
				testCase["input"].([]any)[2].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
