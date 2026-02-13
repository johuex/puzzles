package nearestnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{1, 2, 3, 4, 5}, 6}, "output": 5},
	{"name": "2", "input": []any{[]int{5, 4, 3, 2, 1}, 3}, "output": 3},
}

func TestNearestNumber(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := nearest(testCase["input"].([]any)[0].([]int), testCase["input"].([]any)[1].(int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
