package j

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{
		2,
		14,
		[]Store{
			{
				RPrice: 7,
				OThrs:  9,
				OPrice: 6,
				MTotal: 10,
			},
			{
				RPrice: 7,
				OThrs:  8,
				OPrice: 6,
				MTotal: 10,
			},
		},
	}, "output": []any{
		88,
		[]int{10, 4},
	}},
	{"name": "2", "input": []any{
		1,
		20,
		[]Store{
			{
				RPrice: 1,
				OThrs:  1,
				OPrice: 1,
				MTotal: 1,
			},
		},
	}, "output": []any{
		-1,
		[]int{},
	}},
	{"name": "3", "input": []any{
		2,
		5,
		[]Store{
			{
				RPrice: 15,
				OThrs:  3,
				OPrice: 2,
				MTotal: 4,
			},
			{
				RPrice: 17,
				OThrs:  3,
				OPrice: 3,
				MTotal: 4,
			},
		},
	}, "output": []any{
		15,
		[]int{3, 3},
	}},
	{"name": "4", "input": []any{
		4,
		86,
		[]Store{
			{
				RPrice: 7,
				OThrs:  2,
				OPrice: 3,
				MTotal: 8,
			},
			{
				RPrice: 19,
				OThrs:  34,
				OPrice: 17,
				MTotal: 40,
			},
			{
				RPrice: 13,
				OThrs:  13,
				OPrice: 13,
				MTotal: 13,
			},
			{
				RPrice: 16,
				OThrs:  26,
				OPrice: 12,
				MTotal: 25,
			},
		},
	}, "output": []any{
		15,
		[]int{3, 3},
	}},
}

func TestMasquerade(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res, resArr := masquerade(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].(int),
				testCase["input"].([]any)[2].([]Store),
			)
			assert.Equal(t, testCase["output"].([]any)[0].(int), res)
			assert.Equal(t, testCase["output"].([]any)[1].([]int), resArr)
		})
	}
}
