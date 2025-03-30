package climbingstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": 1, "output": 1},
	{"name": "2", "input": 2, "output": 2},
	{"name": "3", "input": 3, "output": 3},
	{"name": "4", "input": 4, "output": 5},
	{"name": "6", "input": 6, "output": 13},
}

func TestClimbingtairs(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := climbStairs(testCase["input"].(int))
			assert.Equal(t, testCase["output"], res)
		})
	}

}
