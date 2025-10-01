package b

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{1, 2, 2, 10, 10, 10}, "output": "0.500000000000000"},
	{"name": "2", "input": []any{4, 1, 2, 5, 5, 5}, "output": "1.200000000000000"},
	{"name": "3", "input": []any{2, 3, 4, 7, 6, 5}, "output": "1.495238095238095"},
	{"name": "4", "input": []any{1, 6, 3, 7, 6, 5}, "output": "1.271428571428571"},
	{"name": "5", "input": []any{2, 3, 4, 10, 9, 2}, "output": "1.055555555555556"},
}

func TestRoute(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := route(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].(int),
				testCase["input"].([]any)[2].(int),
				testCase["input"].([]any)[3].(int),
				testCase["input"].([]any)[4].(int),
				testCase["input"].([]any)[5].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
