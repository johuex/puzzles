package binaryindexletter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []map[string]any{
	{"name": "1", "input": []any{[]byte{'c', 'f', 'j'}, byte('a')}, "output": byte('c')},
	{"name": "2", "input": []any{[]byte{'c', 'f', 'j'}, byte('c')}, "output": byte('f')},
	{"name": "3", "input": []any{[]byte{'c', 'f', 'j'}, byte('d')}, "output": byte('f')},
	{"name": "4", "input": []any{[]byte{'c', 'f', 'j'}, byte('j')}, "output": byte('c')},
	{"name": "5", "input": []any{[]byte{'x', 'x', 'y', 'y'}, byte('z')}, "output": byte('x')},
}

func TestBinarySearch(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase["name"].(string), func(t *testing.T) {
			res := nextGreatestLetter(
				testCase["input"].([]any)[0].([]byte),
				testCase["input"].([]any)[1].(byte),
			)
			assert.Equal(t, testCase["output"], res)
		})
	}
}
