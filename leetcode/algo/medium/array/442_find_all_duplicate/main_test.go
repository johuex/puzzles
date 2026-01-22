package findduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, "output": []int{2, 3}},
	{"name": "2", "input": []any{[]int{1, 1, 2}}, "output": []int{1}},
	{"name": "3", "input": []any{[]int{1}}, "output": []int{}},
}

func TestFindDuplicate(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := findDuplicates(testCase["input"].([]any)[0].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
