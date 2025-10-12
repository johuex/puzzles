package b

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{"LLBLRRBRL"}, "output": 5},
	{"name": "2", "input": []any{"LRLRLRLRLRLRLR"}, "output": 8},
	{"name": "3", "input": []any{"RBLRBBBLBLBBBRRLLBBRRRRLLBRRBLBRBRRRRLLBRBBLBLRRLRLLRBBLLLRRRLRBBRRRBBRRRRLLBRLBLLBLBRLRLRLBLLLLRBRR"}, "output": 53},
}

func TestRiver(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := river(testCase["input"].([]any)[0].(string))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
