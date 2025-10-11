package g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{1}, "output": 1},
	{"name": "2", "input": []any{3}, "output": 2},
}

func TestStairs(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := stairs(testCase["input"].([]any)[0].(int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
