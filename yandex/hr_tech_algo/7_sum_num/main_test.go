package sumnum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{17, []int{17, 7, 10, 7, 10}}, "output": 4},
	{"name": "2", "input": []any{10, []int{1, 2, 3, 4, 1}}, "output": 2},
}

func TestNearestNumber(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := sumCars(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
