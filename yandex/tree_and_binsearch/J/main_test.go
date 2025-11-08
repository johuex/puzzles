package j

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{
		[]int{6, 14, 70, 1},
		[]int{70, 3, 16, 5},
	},
		"output": 2,
	},
	{"name": "2", "input": []any{
		[]int{2},
		[]int{2},
	},
		"output": 0,
	},
	{"name": "3", "input": []any{
		[]int{3},
		[]int{2},
	},
		"output": -1,
	},
}

func TestLayout(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := interview(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
