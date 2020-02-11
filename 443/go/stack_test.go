package test

import (
	"reflect"
	"testing"
)

type stackT struct {
	push   []interface{}
	expect []interface{} // expect is the expected list of popped elements
}

func TestStack(t *testing.T) {
	tests := []stackT{
		{
			push:   nil,
			expect: nil,
		},
		{
			push:   []interface{}{1, 2, 3},
			expect: []interface{}{3, 2, 1},
		},
		{
			push:   []interface{}{"la", "ti", "da"},
			expect: []interface{}{"da", "ti", "la"},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			stk := stack{elts: queue{}, tmp: queue{}}
			for _, e := range test.push {
				stk.push(e)
			}

			var result []interface{}
			for i := len(test.expect) - 1; i >= 0; i-- {
				result = append(result, stk.pop())
			}

			if !reflect.DeepEqual(result, test.expect) {
				t.Fatalf("expected: %v, got: %v", test.expect, result)
			}

			t.Logf("got expected result: %v", result)
		})
	}
}

// you're a queueT
type queueT struct {
	enqueue []interface{}
	expect  []interface{}
}

func TestQueue(t *testing.T) {
	tests := []queueT{
		{
			enqueue: []interface{}{},
			expect:  nil,
		},
		{
			enqueue: []interface{}{1, 2},
			expect:  []interface{}{1, 2},
		},
		{
			enqueue: []interface{}{"la", "ti", "da"},
			expect:  []interface{}{"la", "ti", "da"},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			q := queue{}
			for _, e := range test.enqueue {
				q.enqueue(e)
			}

			var result []interface{}
			for i := len(test.expect) - 1; i >= 0; i-- {
				result = append(result, q.dequeue())
			}

			if !reflect.DeepEqual(result, test.expect) {
				t.Fatalf("expected: %v, got: %v", test.expect, result)
			}

			t.Logf("got expected result: %v", result)
		})
	}
}
