package minwindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{7, 3, []int{1, 3, 2, 4, 5, 3, 1}}, "output": []int{1, 2, 2, 3, 1}},
	{"name": "2", "input": []any{0, 3, []int{}}, "output": []int{}},
	{"name": "3", "input": []any{7, 0, []int{1, 3, 2, 4, 5, 3, 1}}, "output": []int{}},
}

func TestNearestNumber(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := minWindow(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].(int),
				testCase["input"].([]any)[2].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
