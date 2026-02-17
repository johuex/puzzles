package goodraw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{Wrap([]int{1, 1, 1})}, "output": 2},
	{"name": "2", "input": []any{Wrap([]int{3, 4})}, "output": 3},
	{"name": "3", "input": []any{Wrap([]int{3, 4, 1})}, "output": 4},
	{"name": "4", "input": []any{Wrap([]int{3, 4, 6})}, "output": 7},
	{"name": "5", "input": []any{Wrap([]int{3})}, "output": 0},
	{"name": "6", "input": []any{Wrap([]int{})}, "output": 0},
}

func Wrap(arr []int) [][]int {
	res := make([][]int, len(arr))
	for i := range len(arr) {
		res[i] = []int{i, arr[i]}
	}
	return res
}

func TestGoodRaw(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := goodRow(testCase["input"].([]any)[0].([][]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}

func TestGoodRaw2(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := goodRow2(testCase["input"].([]any)[0].([][]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}

func BenchmarkGoodRaw(b *testing.B) {
	for _, testCase := range testCases {
		b.Run("1_"+testCase["name"].(string), func(t *testing.B) {
			res := goodRow(testCase["input"].([]any)[0].([][]int))
			assert.Equal(t, testCase["output"], res)
		})

		b.Run("2_"+testCase["name"].(string), func(t *testing.B) {
			res := goodRow2(testCase["input"].([]any)[0].([][]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
