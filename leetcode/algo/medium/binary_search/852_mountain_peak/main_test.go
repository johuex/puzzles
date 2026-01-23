package mountainpeak

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []int{0, 1, 0}, "output": 1},
	{"name": "2", "input": []int{-1, 1, 0, -3, 3}, "output": 1},
	{"name": "3", "input": []int{0, 10, 5, 2}, "output": 1},
	{"name": "4", "input": []int{10, 5, 2}, "output": 0},
	{"name": "5", "input": []int{2, 5, 10}, "output": 2},
	{"name": "6", "input": []int{3, 4, 5, 1}, "output": 2},
	{"name": "7", "input": []int{3, 5, 3, 2, 0}, "output": 1},
	{"name": "8", "input": []int{18, 29, 38, 59, 98, 100, 99, 98, 90}, "output": 5},
}

func TestMountainPeak(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := peakIndexInMountainArray(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
