package c

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{"abacaba"}, "output": 15},
	{"name": "2", "input": []any{"aaaaaa"}, "output": 1},
}

func TestTasks(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := passwords(testCase["input"].([]any)[0].(string))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
