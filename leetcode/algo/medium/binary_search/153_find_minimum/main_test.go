package findminimum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []int{3, 4, 5, 1, 2}, "output": 1},
	{"name": "2", "input": []int{4, 5, 6, 7, 0, 1, 2}, "output": 0},
	{"name": "3", "input": []int{11, 13, 15, 17}, "output": 11},
	{"name": "4", "input": []int{17, 11, 13, 15}, "output": 11},
	{"name": "5", "input": []int{15, 17, 11, 13}, "output": 11},
	{"name": "6", "input": []int{13, 15, 17, 11}, "output": 11},

	{"name": "7", "input": []int{11, 12, 13, 15, 17}, "output": 11},
	{"name": "8", "input": []int{17, 11, 12, 13, 15}, "output": 11},
	{"name": "9", "input": []int{15, 17, 11, 12, 13}, "output": 11},
	{"name": "10", "input": []int{13, 15, 17, 11, 12}, "output": 11},
	{"name": "11", "input": []int{12, 13, 15, 17, 11}, "output": 11},

	{"name": "12", "input": []int{1}, "output": 1},
	{"name": "13", "input": []int{1, 2}, "output": 1},
	{"name": "14", "input": []int{3, 4, 5, 6, 1, 2}, "output": 1},
}

func TestMountainPeak(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := findMin(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
