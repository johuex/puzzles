package h

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{1}, "output": 1},
	{"name": "2", "input": []any{2}, "output": 1},
	{"name": "3", "input": []any{4}, "output": 2},
}

func TestWinner(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := winner(testCase["input"].([]any)[0].(int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
