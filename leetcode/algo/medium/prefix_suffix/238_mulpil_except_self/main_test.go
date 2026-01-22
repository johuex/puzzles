package mulpilexceptself

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []int{1, 2, 3, 4}, "output": []int{24, 12, 8, 6}},
	{"name": "2", "input": []int{-1, 1, 0, -3, 3}, "output": []int{0, 0, 9, 0, 0}},
}

func TestMultiplexExceptSelf(t *testing.T) {
	for _, testCase := range testCases {
		t.Run("own"+testCase["name"].(string), func(t *testing.T) {
			res := productExceptSelf(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})

		t.Run("RAM_optmz"+testCase["name"].(string), func(t *testing.T) {
			res := productExceptSelfRamOptimized(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
