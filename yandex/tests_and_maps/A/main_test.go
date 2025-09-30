package a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{2, []int{1, 2}}, "output": 1},
	{"name": "2", "input": []any{3, []int{2, 2, 2}}, "output": 2},
	{"name": "3", "input": []any{11, []int{4, 10, 7, 5, 4, 5, 3, 8, 3, 2, 5}}, "output": 10},
}

func TestMushroomJoy(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := mushroomsJoy(testCase["input"].([]any)[0].(int), testCase["input"].([]any)[1].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
