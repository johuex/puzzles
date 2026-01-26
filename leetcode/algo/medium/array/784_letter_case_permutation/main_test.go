package lettercasepermutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{"a1b2"}, "output": []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
	{"name": "2", "input": []any{"3z4"}, "output": []string{"3z4", "3Z4"}},
}

func TestLetterCasePermutation(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := letterCasePermutation(testCase["input"].([]any)[0].(string))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
