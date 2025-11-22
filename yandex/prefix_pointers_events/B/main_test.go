package b

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			[][]int{
				{0, 24},
				{100, 35},
				{150, 50},
				{200, 75},
				{250, 150},
			},
			[]int{
				107,
				143,
				152,
				170,
				150,
			},
		},
		"output": []int{
			3745,
			5005,
			7600,
			8500,
			5250,
		},
	},
}

func TestRiver(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := taxCalculator(
				testCase["input"].([]any)[0].([][]int),
				testCase["input"].([]any)[1].([]int),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
