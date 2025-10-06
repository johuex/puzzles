package f

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{4, 3, []string{
			"+-+",
			"??-",
			"?-?",
			"++?",
		}},
		"output": 5,
	},
	{
		"name": "2",
		"input": []any{6, 10, []string{
			"??+++-?-?-",
			"-??+???--+",
			"?-+?+-?+--",
			"??????--?+",
			"++--?--+-?",
			"?-?+++?+-?",
		}},
		"output": 12,
	},
	{
		"name": "3",
		"input": []any{6, 3, []string{
			"-++",
			"?+-",
			"+-+",
			"???",
			"???",
			"?+?",
		}},
		"output": 7,
	},
}

func TestTasks(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := pcq(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].(int),
				testCase["input"].([]any)[2].([]string),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
