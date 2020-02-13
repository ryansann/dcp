package test_test

import (
	"fmt"
	"reflect"
	"testing"

	s "github.com/ryansann/dcp/436/go"
)

type opcode uint

const (
	push opcode = iota
	pop
)

var opcodeNames = map[opcode]string{
	push: "push",
	pop:  "pop",
}

type op struct {
	typ       opcode
	n         uint
	value     interface{}
	expect    interface{}
	expectErr bool
}

type multiStackT struct {
	count uint
	ops   []op
}

func TestMultiStacker(t *testing.T) {
	tests := []multiStackT{
		{
			count: 1,
			ops: []op{
				{
					typ:   push,
					n:     0,
					value: 1,
				},
				{
					typ:   push,
					n:     0,
					value: 2,
				},
				{
					typ:    pop,
					n:      0,
					expect: 2,
				},
				{
					typ:   push,
					n:     0,
					value: 2,
				},
				{
					typ:    pop,
					n:      0,
					expect: 2,
				},
				{
					typ:    pop,
					n:      0,
					expect: 1,
				},
			},
		},
		{
			count: 2,
			ops: []op{
				{
					typ:   push,
					n:     0,
					value: 1,
				},
				{
					typ:   push,
					n:     0,
					value: 2,
				},
				{
					typ:    pop,
					n:      0,
					expect: 2,
				},
				{
					typ:   push,
					n:     1,
					value: 2,
				},
				{
					typ:    pop,
					n:      1,
					expect: 2,
				},
				{
					typ:   push,
					n:     1,
					value: 3,
				},
				{
					typ:   push,
					n:     1,
					value: 4,
				},
				{
					typ:    pop,
					n:      1,
					expect: 4,
				},
				{
					typ:    pop,
					n:      0,
					expect: 1,
				},
			},
		},
		{
			count: 3,
			ops: []op{
				{
					typ:   push,
					n:     0,
					value: 1,
				},
				{
					typ:   push,
					n:     1,
					value: 1,
				},
				{
					typ:   push,
					n:     2,
					value: 1,
				},
				{
					typ:    pop,
					n:      0,
					expect: 1,
				},
				{
					typ:    pop,
					n:      1,
					expect: 1,
				},
				{
					typ:    pop,
					n:      2,
					expect: 1,
				},
			},
		},
		{
			count: 3,
			ops: []op{
				{
					typ:   push,
					n:     0,
					value: "world",
				},
				{
					typ:   push,
					n:     1,
					value: "world",
				},
				{
					typ:   push,
					n:     0,
					value: "hello",
				},
				{
					typ:   push,
					n:     1,
					value: "hello",
				},
				{
					typ:   push,
					n:     2,
					value: "world",
				},
				{
					typ:   push,
					n:     2,
					value: "hello",
				},
				{
					typ:    pop,
					n:      0,
					expect: "hello",
				},
				{
					typ:    pop,
					n:      1,
					expect: "hello",
				},
				{
					typ:    pop,
					n:      2,
					expect: "hello",
				},
				{
					typ:    pop,
					n:      0,
					expect: "world",
				},
				{
					typ:    pop,
					n:      2,
					expect: "world",
				},
				{
					typ:    pop,
					n:      1,
					expect: "world",
				},
			},
		},
		{
			count: 5,
			ops: []op{
				{
					typ:       push, // n out of bounds
					n:         5,
					value:     1,
					expectErr: true,
				},
				{
					typ:       pop, // n out of bounds
					n:         5,
					expectErr: true,
				},
				{
					typ:       push, // n out of bounds
					n:         6,
					value:     1,
					expectErr: true,
				},
				{
					typ:       pop, // n out of bounds
					n:         6,
					expectErr: true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("n=%v", test.count), func(t *testing.T) {
			ms, err := s.NewMultiStacker(test.count)
			if err != nil {
				t.Fatalf("could not create multistacker: %v", err)
			}

			for opCount, operation := range test.ops {
				stmt := fmt.Sprintf("operation %v: %s", opCount, opcodeNames[operation.typ])
				switch operation.typ {
				case push:
					stmt = fmt.Sprintf("%s value=%v n=%v", stmt, operation.value, operation.n)
					t.Log(stmt)

					err := ms.Push(operation.value, operation.n)
					if err != nil && !operation.expectErr { // got an unexpected error
						t.Fatalf("got unexpected error: %v", err)
					}
				case pop:
					stmt = fmt.Sprintf("%s n=%v", stmt, operation.n)
					t.Log(stmt)

					v, err := ms.Pop(operation.n)
					if err != nil && !operation.expectErr { // got an unexpected error
						t.Fatalf("got unexpected error: %v", err)
					}

					if !reflect.DeepEqual(v, operation.expect) {
						t.Fatalf("expected: %v, got: %v", operation.expect, v)
					}
				}

				opCount++
			}
		})
	}
}
