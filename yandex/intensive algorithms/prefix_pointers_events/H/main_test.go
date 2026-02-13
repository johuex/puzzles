package h

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			[]int{4, 2, 2, 4},
		},
		"output": 14,
	},
	{
		"name": "2",
		"input": []any{
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		"output": 0,
	},
}

func TestBoss(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := boss(
				testCase["input"].([]any)[0].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
