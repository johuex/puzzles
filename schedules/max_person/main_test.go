package main

import (
	"testing"
)

type Case struct {
	input    []string
	expected int
}

var cases = []Case{
	{[]string{"5, 8", "2, 4", "3, 9"}, 2},
	{[]string{"4, 7", "1, 3"}, 1},
}

func TestBench(t *testing.T) {
	var got, want int
	for _, case_ := range cases {
		got = MaxCustomers(case_.input)
		want = case_.expected

		if got != want {
			t.Errorf("got %d, wanted %d", got, want)
		}
	}

}
