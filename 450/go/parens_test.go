package test

import (
	"fmt"
	"testing"
)

type isBalancedT struct {
	input  string
	expect bool
}

type isBalancedFn func(s string) bool

func Test(t *testing.T) {
	tests := []isBalancedT{
		{
			input:  "",
			expect: true,
		},
		{
			input:  "()",
			expect: true,
		},
		{
			input:  ")(",
			expect: false,
		},
		{
			input:  "(())",
			expect: true,
		},
		{
			input:  "()()",
			expect: true,
		},
		{
			input:  "(()())",
			expect: true,
		},
		{
			input:  "(()",
			expect: false,
		},
		{
			input:  "(*)",
			expect: true,
		},
		{
			input:  "((*)",
			expect: true,
		},
		{
			input:  "()*)",
			expect: true,
		},
		{
			input:  "(()*",
			expect: true,
		},
		{
			input:  ")*(",
			expect: false,
		},
		{
			input:  "()********",
			expect: true,
		},
		{
			input:  "(*********)",
			expect: true,
		},
		{
			input:  "************",
			expect: true,
		},
		{
			input:  "()*()**()***()((((((*))))))",
			expect: true,
		},
		{
			input:  "()*()**()***()((((((*)))))))",
			expect: true,
		},
		{
			input:  "()*()**()***()(((((())))))))))))))))",
			expect: false,
		},
	}

	testFns := map[string]isBalancedFn{
		"brute-force": isBalanced,
		"greedy":      isBalancedGreedy,
	}

	for _, test := range tests {
		for k, fn := range testFns {
			t.Run(fmt.Sprintf("fn=%s|input=%s,expect=%v", k, test.input, test.expect), func(t *testing.T) {
				result := fn(test.input)
				if result != test.expect {
					t.Fatalf("expected: %v, got: %v", test.expect, result)
				}
			})
		}
	}
}
