package test

import (
	"fmt"
	"testing"
)

type powerT struct {
	input  uint32
	expect bool
}

type isPowerFn func(n uint32) bool

func Test(t *testing.T) {
	tests := []powerT{
		{
			input:  0,
			expect: false,
		},
		{
			input:  2,
			expect: false,
		},
		{
			input:  4,
			expect: true,
		},
		{
			input:  12,
			expect: false,
		},
		{
			input:  16,
			expect: true,
		},
		{
			input:  9,
			expect: false,
		},
		{
			input:  64,
			expect: true,
		},
		{
			input:  128,
			expect: false,
		},
		{
			input:  256,
			expect: true,
		},
		{
			input:  257,
			expect: false,
		},
		{
			input:  512,
			expect: false,
		},
		{
			input:  1024, // 4^5
			expect: true,
		},
		{
			input:  16777216, // 4^12
			expect: true,
		},
	}

	testFns := []isPowerFn{
		isPowerOfFour,
		isPowerOfFourEasy,
	}

	for _, test := range tests {
		for i, fn := range testFns {
			t.Run(fmt.Sprintf("testFns[%v](%v)", i, test.input), func(t *testing.T) {
				run(t, fn, test)
			})
		}
	}
}

// run executes a test with the provided isPowerFn
func run(t *testing.T, fn isPowerFn, test powerT) {
	if result := fn(test.input); result != test.expect {
		t.Fatalf("expected: %v, got: %v", test.expect, result)
	}
}
