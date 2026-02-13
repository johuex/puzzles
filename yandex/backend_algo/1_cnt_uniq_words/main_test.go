package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{bufio.NewReader(strings.NewReader(`She sells sea shells on the sea shore;
The shells that she sells are sea shells I'm sure.
So if she sells sea shells on the sea shore,
I'm sure that the shells are sea shore shells.`))}, "output": 19},
	{"name": "2", "input": []any{bufio.NewReader(strings.NewReader("AA aa Aa aA"))}, "output": 4},
	{"name": "3", "input": []any{bufio.NewReader(strings.NewReader(`a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a a
`))}, "output": 1},
}

func TestNearestNumber(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := countWords(testCase["input"].([]any)[0].(*bufio.Reader))
			assert.Equal(t, testCase["output"], res)
		})
	}
}
