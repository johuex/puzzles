package wordsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCCED",
		},
		"output": true,
	},
	{
		"name": "2",
		"input": []any{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SEE",
		},
		"output": true,
	},
	{
		"name": "3",
		"input": []any{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCB",
		},
		"output": false,
	},
	{
		"name": "4", // TL check
		"input": []any{
			[][]byte{
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'B'},
				{'A', 'A', 'A', 'A', 'B', 'A'},
			},
			"AAAAAAAAAAAAABB",
		},
		"output": false,
	},
	{
		"name": "5",
		"input": []any{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'E', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"ABCESEEEFS",
		},
		"output": true,
	},
}

func TestFindWord(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := exist(
				testCase["input"].([]any)[0].([][]byte),
				testCase["input"].([]any)[1].(string),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}

func TestFindWordFaster(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := existFaster(
				testCase["input"].([]any)[0].([][]byte),
				testCase["input"].([]any)[1].(string),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
