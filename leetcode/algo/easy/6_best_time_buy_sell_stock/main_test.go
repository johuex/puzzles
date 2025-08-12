package beststock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []int{7, 2, 5, 1, 3, 6}, "output": 5},
	{"name": "2", "input": []int{7, 2, 8, 1, 6, 4}, "output": 6},
	{"name": "3", "input": []int{7, 6, 4, 3, 1}, "output": 0},
}

func TestMaxProfit(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := maxProfit(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}

}

func TestMaxProfitOpt(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := maxProfitOpt(testCase["input"].([]int))
			assert.Equal(t, testCase["output"], res)
		})
	}

}
