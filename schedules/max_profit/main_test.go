package main

import (
	"testing"
)

type Case struct {
	input    []string
	expected int
}

var cases = []Case{
	{[]string{"A, 4, 2", "B, 3, 10", "C, 2, 8", "D, 4, 15"}, 6},
}

func TestBench(t *testing.T) {
	var got, want int
	for _, case_ := range cases {
		got = MaxProfit(case_.input)
		want = case_.expected

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}
}
