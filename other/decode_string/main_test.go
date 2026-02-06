package decodestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name":   "1",
		"input":  []any{"3[a]2[bc]"},
		"output": "aaabcbc",
	},
	{
		"name":   "2",
		"input":  []any{"3[a2[c]]"},
		"output": "accaccacc",
	},
	{
		"name":   "3",
		"input":  []any{"2[abc]3[cd]ef"},
		"output": "abcabccdcdcdef",
	},
	{
		"name":   "4",
		"input":  []any{"3[at2[c2[d]2[i]ef]]"},
		"output": "atcddiiefcddiiefatcddiiefcddiiefatcddiiefcddiief",
	},
}

func TestDecodeString(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := DecodeString(testCase["input"].([]any)[0].(string))
			assert.Equal(t, testCase["output"], res)
		})
	}

}
