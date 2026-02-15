package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[][]int{
		{0, 5},
		{-3, 2},
		{7, 10},
	},
		[]int{1, 6}}, "output": []int{2, 0}},
}

func TestNearestNumber(t *testing.T) {
	for _, testCase := range testCases {
		t.Run("1_"+testCase["name"].(string), func(t *testing.T) {
			res := cntPoints(testCase["input"].([]any)[0].([][]int), testCase["input"].([]any)[1].([]int))
			assert.Equal(t, testCase["output"], res)
		})
		t.Run("2_"+testCase["name"].(string), func(t *testing.T) {
			res := cntPoints(testCase["input"].([]any)[0].([][]int), testCase["input"].([]any)[1].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}

func BenchmarkNearestNumber(b *testing.B) {
	for _, testCase := range testCases {
		b.Run("1_"+testCase["name"].(string), func(t *testing.B) {
			res := cntPoints(testCase["input"].([]any)[0].([][]int), testCase["input"].([]any)[1].([]int))
			assert.Equal(t, testCase["output"], res)
		})
		b.Run("2_"+testCase["name"].(string), func(t *testing.B) {
			res := cntPoints(testCase["input"].([]any)[0].([][]int), testCase["input"].([]any)[1].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
