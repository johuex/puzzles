package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: upd test input
var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			[][]int{
				{-4, -1, 1},
				{13, 6, 3},
				{-7, -6, 1},
			},
			[]int{1, 5},
			0,
		},
		"output": []float64{4.333333333333333, 5.000000000},
	},
	{
		"name": "2",
		"input": []any{
			[][]int{
				{4, 2, 1},
				{-11, -8, 2},
			},
			[]int{2, 6},
			0,
		},
		"output": []float64{5.500000000, 6.000000000},
	},
}

func TestFriends(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := railwayCrossing(
				testCase["input"].([]any)[0].([][]int),
				testCase["input"].([]any)[1].([]int),
				testCase["input"].([]any)[2].(int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
