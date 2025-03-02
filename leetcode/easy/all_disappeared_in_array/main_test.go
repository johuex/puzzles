package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDisappearedNumbers(t *testing.T) {
	testCases := []map[string]any{
		{"name": "1", "input": []int{2, 2, 3, 4}, "output": []int{1}},
		{"name": "2,4", "input": []int{1, 1, 3, 3, 5}, "output": []int{2, 4}},
		{"name": "--", "input": []int{1, 2, 3, 4}, "output": []int{}},
	}

	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := findDisappearedNumbers(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}

}
