package test

import (
	"fmt"
	"testing"
)

type isBalancedT struct {
	input  string
	expect bool
}

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
		{
			input:  "la",
			expect: false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input=%s,expect=%v", test.input, test.expect), func(t *testing.T) {
			result := isBalanced(test.input)
			if result != test.expect {
				t.Fatalf("expected: %v, got: %v", test.expect, result)
			}
		})
	}
}
