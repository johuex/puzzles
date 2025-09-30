package e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{1, 10}, "output": 44},
	{"name": "2", "input": []any{5, 1}, "output": 10},
	{"name": "3", "input": []any{0, 10000000000}, "output": 0},
	{"name": "4", "input": []any{25, 0}, "output": 25},
	{"name": "5", "input": []any{35, 1}, "output": 40},
	{"name": "6", "input": []any{51, 926907757}, "output": 4634538832},
	{"name": "7", "input": []any{51, 926907756}, "output": 4634538826},
	{"name": "8", "input": []any{51, 926907758}, "output": 4634538834},
	{"name": "9", "input": []any{51, 926907759}, "output": 4634538838},
	{"name": "10", "input": []any{2, 10}, "output": 48},
	{"name": "11", "input": []any{3, 10}, "output": 52},
	{"name": "12", "input": []any{4, 10}, "output": 56},
	{"name": "13", "input": []any{6, 10}, "output": 54},
	{"name": "14", "input": []any{7, 10}, "output": 58},
	{"name": "15", "input": []any{8, 10}, "output": 62},
	{"name": "16", "input": []any{9, 10}, "output": 66},
	{"name": "17", "input": []any{52, 926907757}, "output": 4634538834},
	{"name": "18", "input": []any{52, 926907756}, "output": 4634538832},
	{"name": "19", "input": []any{52, 926907758}, "output": 4634538838},
	{"name": "20", "input": []any{52, 926907759}, "output": 4634538846},
}

func TestTasks(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := sumNum(testCase["input"].([]any)[0].(int), testCase["input"].([]any)[1].(int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
