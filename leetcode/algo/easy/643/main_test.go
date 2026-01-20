package findmaxaverage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{1, 12, -5, -6, 50, 3}, 4}, "output": 12.75000},
	{"name": "2", "input": []any{[]int{5}, 1}, "output": float64(5)},
}

func TestBinarySearch(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := findMaxAverage(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
