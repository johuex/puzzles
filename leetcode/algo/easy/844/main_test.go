package backspacecompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{"ab#c", "ad#c"}, "output": true},
	{"name": "2", "input": []any{"ab##", "c#d#"}, "output": true},
	{"name": "3", "input": []any{"a#c", "b"}, "output": false},
	{"name": "4", "input": []any{"a##c", "#a#c"}, "output": true},
	{"name": "5", "input": []any{"isfcow#", "isfcog#w#"}, "output": true},
}

func TestSortedSquares(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := backspaceCompare(
				testCase["input"].([]any)[0].(string),
				testCase["input"].([]any)[1].(string),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
