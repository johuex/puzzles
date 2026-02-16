package f

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{
		5,
		[]string{
			"W.W",
			"C.C",
			"WW.",
			"CC.",
			"CWW",
		},
	}, "output": 3},
	{"name": "2", "input": []any{
		4,
		[]string{
			"W.W",
			"CWC",
			"W.W",
			"CWW",
		},
	}, "output": 2},
	{"name": "3", "input": []any{
		5,
		[]string{
			"W.W",
			"C.C",
			"WWW",
			"CC.",
			"CWW",
		},
	}, "output": 0},
	{"name": "4", "input": []any{
		4,
		[]string{
			"WWW",
			"CWC",
			"W.W",
			"CWW",
		},
	}, "output": 0},
	{"name": "5", "input": []any{
		7,
		[]string{
			"W.W",
			"CW.",
			"WWC",
			"CW.",
			"CW.",
			"C..",
			".WW",
		},
	}, "output": 1},
	{"name": "6", "input": []any{
		6,
		[]string{
			"WW.",
			"CWW",
			"WW.",
			"CW.",
			"CWC",
			"WWC",
		},
	}, "output": 0},
	{"name": "7", "input": []any{
		3,
		[]string{
			"...",
			"...",
			"...",
		},
	}, "output": 0},
	{"name": "7", "input": []any{
		3,
		[]string{
			"CCC",
			"CCC",
			"CCC",
		},
	}, "output": 3},
	{"name": "8", "input": []any{
		3,
		[]string{
			".C.",
			".WC",
			"...",
		},
	}, "output": 2},
	{"name": "9", "input": []any{
		17,
		[]string{
			"WW.",
			"..C",
			".C.",
			".CC",
			".WC",
			"WW.",
			".WW",
			"CW.",
			".CW",
			"CWC",
			".WC",
			"CW.",
			"CWC",
			"CWC",
			".W.",
			"CCW",
			"CWW",
		},
	}, "output": 4},
	{"name": "10", "input": []any{
		29,
		[]string{
			"..C",
			"C..",
			"WWW",
			"WWC",
			"...",
			".WC",
			".C.",
			"WW.",
			"..C",
			"C.W",
			".W.",
			"W.W",
			".CW",
			"CCC",
			"...",
			"C.C",
			"C..",
			"WCC",
			"CWW",
			"C..",
			"...",
			".CC",
			".WC",
			"WWW",
			"CCW",
			".WW",
			"C..",
			"WC.",
			"W.C",
		},
	}, "output": 1},
}

func TestCoinMax(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := coinMax(
				testCase["input"].([]any)[0].(int),
				testCase["input"].([]any)[1].([]string),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
