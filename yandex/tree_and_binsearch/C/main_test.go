package c

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{10, 7, [][]float64{
		{4, 3},
		{3, 2},
		{4, 2},
	}}, "output": "1.400000"},
	{"name": "2", "input": []any{10, 1, [][]float64{
		{2, 1},
		{3, 2},
	}}, "output": "0.333333"},
}

func TestLayout(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := layout(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].(int),
				testCase["input"].([]any)[2].([][]float64),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
