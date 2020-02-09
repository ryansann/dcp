package test

import (
	"reflect"
	"testing"
)

type closureT struct {
	input  graph
	expect closure
}

func TestTransitiveClosure(t *testing.T) {
	tests := []closureT{
		{
			input:  nil,
			expect: nil,
		},
		{
			input: [][]int{
				[]int{0},
			},
			expect: [][]bool{
				[]bool{true},
			},
		},
		{
			input: [][]int{
				[]int{},
			},
			expect: [][]bool{
				[]bool{true},
			},
		},
		{
			input: [][]int{
				[]int{0, 1, 3},
				[]int{1, 2},
				[]int{2},
				[]int{3},
			},
			expect: [][]bool{
				[]bool{true, true, true, true},
				[]bool{false, true, true, false},
				[]bool{false, false, true, false},
				[]bool{false, false, false, true},
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			c := transitiveClosure(test.input)
			if !reflect.DeepEqual(test.expect, c) {
				t.Fatalf("expected: %v, got: %v", test.expect, c)
			}
			t.Logf("expected: %v, got: %v", test.expect, c)
		})
	}
}
