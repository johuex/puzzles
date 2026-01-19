package issubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{"abc", "ahbgdc"}, "output": true},
	{"name": "2", "input": []any{"axc", "ahbgdc"}, "output": false},
}

func TestIsSubsequence(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := isSubsequence(
				testCase["input"].([]any)[0].(string),
				testCase["input"].([]any)[1].(string),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
