package test

import (
	"fmt"
	"testing"
)

type matchT struct {
	n      string
	k      string
	expect int
}

func TestMatch(t *testing.T) {
	tests := []matchT{
		{
			n:      "",
			k:      "",
			expect: -1,
		},
		{
			n:      "a",
			k:      "a",
			expect: 0,
		},
		{
			n:      "hello",
			k:      "ll",
			expect: 2,
		},
		{
			n:      "asdasfksl",
			k:      "asf",
			expect: 3,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("n=%s|k=%s", test.n, test.k), func(t *testing.T) {
			maxIterations := len([]rune(test.n)) * len([]rune(test.k))
			t.Logf("maxIterations: %v", maxIterations)
			nPos, iterations := match(test.n, test.k)
			t.Logf("nPos: %v, iterations: %v", nPos, iterations)
			if nPos != test.expect {
				t.Errorf("expected: %v, got: %v", test.expect, nPos)
			}
			if iterations > maxIterations {
				t.Errorf("expected < %v iterations, got: %v", maxIterations, iterations)
			}
		})
	}
}
