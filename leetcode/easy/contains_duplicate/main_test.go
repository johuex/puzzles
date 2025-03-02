package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicates(t *testing.T) {
	testCases := []map[string]any{
		{"name": "all unique", "input": []int{1, 2, 3}, "output": false},
		{"name": "two duplicates", "input": []int{1, 2, 1}, "output": true},
		{"name": "empty", "input": []int{}, "output": false},
		{"name": "only one elem", "input": []int{1}, "output": false},
	}

	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := containsDuplicate(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}

}
