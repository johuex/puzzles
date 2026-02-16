package a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{4}, "output": 7},
	{"name": "2", "input": []any{5}, "output": 13},
	{"name": "3", "input": []any{1}, "output": 1},
	{"name": "4", "input": []any{2}, "output": 2},
	{"name": "5", "input": []any{3}, "output": 4},
}

func TestMushroomJoy(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := ballJump(testCase["input"].([]any)[0].(int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
