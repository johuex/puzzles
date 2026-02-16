package niceraw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{2, "abcaz"}, "output": 4},
	{"name": "2", "input": []any{2, "helto"}, "output": 3},
	{"name": "3", "input": []any{0, "helto"}, "output": 1},
	{"name": "4", "input": []any{1, ""}, "output": 0},
}

func TestNiceRaw(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := niceRaw(testCase["input"].([]any)[0].(int), testCase["input"].([]any)[1].(string))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
