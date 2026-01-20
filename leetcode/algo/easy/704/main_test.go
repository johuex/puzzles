package binaryindex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{-1, 0, 3, 5, 9, 12}, 9}, "output": 4},
	{"name": "2", "input": []any{[]int{-1, 0, 3, 5, 9, 12}, 2}, "output": -1},
	{"name": "3", "input": []any{[]int{-1, 0, 3, 5, 9, 12}, 13}, "output": -1},
}

func TestBinarySearch(t *testing.T) {
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
