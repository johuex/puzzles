package a

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{
		"name": "1",
		"input": []any{
			[]string{
				"06:45-10:20",
				"07:36-11:26",
				"19:00-22:35",
				"20:08-23:58",
			},
			[]string{
				"06:35-10:10",
				"07:15-11:10",
				"11:00-14:48",
				"14:00-17:48",
				"15:40-19:28",
				"18:35-22:23",
				"20:20-23:55",
			},
		},
		"output": 7,
	},
	{
		"name": "2",
		"input": []any{
			[]string{
				"10:00-12:00",
				"15:00-17:00",
			},
			[]string{
				"12:30-14:30",
				"17:30-19:30",
			},
		},
		"output": 1,
	},
	{
		"name": "3",
		"input": []any{
			[]string{
				"10:10-10:11",
				"10:10-10:11",
			},
			[]string{
				"10:11-10:12",
				"10:11-10:12",
			},
		},
		"output": 2,
	},
}

func TestFriends(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := schedule(
				testCase["input"].([]any)[0].([]string),
				testCase["input"].([]any)[1].([]string),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
