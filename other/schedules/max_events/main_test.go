package main

import (
	"testing"
)

type Case struct {
	input    []string
	expected []int
}

var cases = []Case{
	{[]string{"1, 4", "2, 5", "5, 8"}, []int{1, 3}},
	{[]string{"1, 3", "1, 2", "3, 6"}, []int{2, 3}},
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestBench(t *testing.T) {
	var got, want []int
	for _, case_ := range cases {
		got = MaxEvent(case_.input)
		want = case_.expected

		if !Equal(got, want) {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
}
