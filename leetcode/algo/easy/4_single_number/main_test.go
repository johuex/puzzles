package singlenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []int{2, 2, 1}, "output": 1},
	{"name": "2", "input": []int{4, 1, 2, 1, 2}, "output": 4},
	{"name": "3", "input": []int{1}, "output": 1},
}

func TestSingleNumber(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := singleNumber(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}

func TestSingleNumberXOR(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := singleNumberXOR(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
