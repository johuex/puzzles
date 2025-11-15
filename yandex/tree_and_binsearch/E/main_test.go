package e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{
		[]int{-1, 0, 0},
		[]int{2, 1, 3},
	},
		"output": 6,
	},
	{"name": "2", "input": []any{
		[]int{-1, 0, 0, 2, 2},
		[]int{0, 1, 2, 1, 3},
	},
		"output": 7,
	},
	{"name": "3", "input": []any{
		[]int{-1, 0, 1, 1},
		[]int{0, 5, 1, 1},
	},
		"output": 5,
	},
	{"name": "4", "input": []any{
		[]int{-1, 0, 0, 0, 1, 0, 1, 0, 3, 0},
		[]int{-58696, -84845, 4890, -80991, -6177, -71108, -36794, 81952, 28859, -42781},
	},
		"output": 558472,
	},
}

func TestParent(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := taxes(
				testCase["input"].([]any)[0].([]int),
				testCase["input"].([]any)[1].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
