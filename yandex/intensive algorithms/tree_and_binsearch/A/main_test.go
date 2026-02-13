package a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{1, 1, 1}, "output": 2},
	{"name": "2", "input": []any{2, 3, 12}, "output": 6},
	{"name": "3", "input": []any{1, 1, 2}, "output": -1},
	{"name": "4", "input": []any{0, 2, 3}, "output": 3},
	{"name": "5", "input": []any{2, 0, 3}, "output": 3},
	{"name": "6", "input": []any{0, 0, 3}, "output": -1},
	{"name": "7", "input": []any{0, 0, 1}, "output": -1},
	{"name": "8", "input": []any{566, 239, 664660}, "output": 1234},
	{"name": "9", "input": []any{5984, 6045, 8772}, "output": 6113},
}

func TestShelf(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := shelf(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].(int),
				testCase["input"].([]any)[2].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
