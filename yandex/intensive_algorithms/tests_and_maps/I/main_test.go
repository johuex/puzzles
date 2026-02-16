package i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{8, 7, 3, 4}, "output": 19},
	{"name": "2", "input": []any{3, 4, 4, 3}, "output": 1},
}

func TestRoute(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := route(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].(int),
				testCase["input"].([]any)[2].(int),
				testCase["input"].([]any)[3].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
