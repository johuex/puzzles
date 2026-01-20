package majorityelement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{3, 2, 3}}, "output": 3},
	{"name": "2", "input": []any{[]int{2, 2, 1, 1, 1, 2, 2}}, "output": 2},
	{"name": "3", "input": []any{[]int{1}}, "output": 1},
}

func TestBinarySearch(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := majorityElement(
				testCase["input"].([]any)[0].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
