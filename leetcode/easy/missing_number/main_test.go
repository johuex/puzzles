package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	testCases := []map[string]any{
		{"name": "1", "input": []int{0, 2, 3}, "output": 1},
		{"name": "2", "input": []int{4, 0, 1, 3}, "output": 2},
		{"name": "0", "input": []int{3, 4, 1, 2}, "output": 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := missingNumber(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}

}
