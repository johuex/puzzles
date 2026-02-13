package d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{3, []int{1, 1, 1, 2, 2}}, "output": []int{2, 1, 2}},
	{"name": "2", "input": []any{4, []int{8, 8, 8, 8, 8, 8, 8, 8, 2, 1}}, "output": []int{2, 1, 8, 8}},
}

func TestTasks(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := tasks(testCase["input"].([]any)[0].(int), testCase["input"].([]any)[1].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
