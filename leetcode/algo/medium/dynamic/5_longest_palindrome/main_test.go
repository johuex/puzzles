package longestpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name":   "1",
		"input":  []any{"babad"},
		"output": "bab",
	},
	{
		"name":   "2",
		"input":  []any{"cbbd"},
		"output": "bb",
	},
	{
		"name":   "3",
		"input":  []any{"ac"},
		"output": "a",
	},
	{
		"name":   "4",
		"input":  []any{"abcbabcdd"},
		"output": "abcba",
	},
	{
		"name":   "5",
		"input":  []any{"babadada"},
		"output": "adada",
	},
}

func TestLongestPalindrome(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := longestPalindrome(testCase["input"].([]any)[0].(string))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
