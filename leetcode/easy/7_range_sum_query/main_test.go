package range_sum_query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var initArr = []int{-2, 0, 3, -5, 2, -1}

var testCases = []map[string]any{
	{"name": "1", "input": []int{0, 2}, "output": 1},
	{"name": "2", "input": []int{2, 5}, "output": -1},
	{"name": "3", "input": []int{0, 5}, "output": -3},
}

func TestRangeSumStupid(t *testing.T) {
	obj := Constructor(initArr)
	for _, testCase := range testCases {
		left := testCase["input"].([]int)[0]
		right := testCase["input"].([]int)[1]
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := obj.SumRange(left, right)
			assert.Equal(t, testCase["output"], res)
		})
	}

}

func TestRangeSumBest(t *testing.T) {
	obj := Constructor1(initArr)
	for _, testCase := range testCases {
		left := testCase["input"].([]int)[0]
		right := testCase["input"].([]int)[1]
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := obj.SumRange(left, right)
			assert.Equal(t, testCase["output"], res)
		})
	}

}
