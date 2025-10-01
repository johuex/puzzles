package g

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{5, 6, []string{
			".OXOXO",
			"OXX.O.",
			"XXOOXX",
			"..O..O",
			"OOXX.X",
		}},
		"output": "Yes",
	},
	{
		"name": "2",
		"input": []any{2, 6, []string{
			"XX....",
			".XXXXX",
		}},
		"output": "Yes",
	},
	{
		"name": "3",
		"input": []any{6, 6, []string{
			"XX....",
			"OXX.XX",
			".O...O",
			"..O...",
			"...O..",
			"...OO.",
		}},
		"output": "Yes",
	},
	{
		"name": "4",
		"input": []any{2, 6, []string{
			"XX.Y..",
			"YXXXXX",
		}},
		"output": "Yes",
	},
}

func TestTasks(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := fiveRow(testCase["input"].([]any)[2].([]string))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
