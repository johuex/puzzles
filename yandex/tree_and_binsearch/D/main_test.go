package d

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{2, []int{1, 6, 3}}, "output": []int{2, 3}},
	{"name": "2", "input": []any{3, []int{2, 3, 4, 5}}, "output": []int{4, 1}},
	{"name": "3", "input": []any{2690, []int{22434, 22385, 22450, 22378, 60196906, 60196806, 22473, 60196871, 60196824, 22451}}, "output": []int{9, 4}},
}

func TestShelf(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := exchange(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
