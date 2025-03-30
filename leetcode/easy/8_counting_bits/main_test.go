package counting_bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "2", "input": 2, "output": []int{0, 1, 1}},
	{"name": "5", "input": 5, "output": []int{0, 1, 1, 2, 1, 2}},
	{"name": "10", "input": 10, "output": []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2}},
}

func TestCountBits(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := countBits(testCase["input"].(int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}

func TestCountBitsLow(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := countBitsLow(testCase["input"].(int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
