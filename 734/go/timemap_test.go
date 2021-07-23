package test

import (
	"fmt"
	"testing"
)

type kvt struct {
	key  string
	val  string
	time int
}

type timeMapTest struct {
	set []kvt
	get []kvt
}

func TestTimeMap(t *testing.T) {
	tests := []timeMapTest{
		{
			set: []kvt{
				{
					key:  "a",
					val:  "b",
					time: 1,
				},
				{
					key:  "a",
					val:  "c",
					time: 3,
				},
			},
			get: []kvt{
				{
					key:  "a",
					time: 2,
					val:  "b",
				},
			},
		},
		{
			set: []kvt{
				{
					key:  "a",
					val:  "b",
					time: -1,
				},
				{
					key:  "a",
					val:  "c",
					time: 1,
				},
			},
			get: []kvt{
				{
					key:  "a",
					time: 0,
					val:  "b",
				},
			},
		},
		{
			set: []kvt{
				{
					key:  "a",
					val:  "b",
					time: 0,
				},
			},
			get: []kvt{
				{
					key:  "a",
					time: 0,
					val:  "b",
				},
			},
		},
		{
			set: []kvt{
				{
					key:  "a",
					val:  "b",
					time: 0,
				},
			},
			get: []kvt{
				{
					key:  "a",
					time: 1,
					val:  "b",
				},
			},
		},
		{
			set: []kvt{
				{
					key:  "a",
					val:  "b",
					time: 1,
				},
				{
					key:  "a",
					val:  "c",
					time: 3,
				},
				{
					key: "a",
					val: "d",
					time: 5,
				},
			},
			get: []kvt{
				{
					key:  "a",
					time: 7,
					val:  "d",
				},
			},
		},
	}

	for i, tst := range tests {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			// create map and perform set operations
			m := newTimeBasedMap(1)
			for _, s := range tst.set {
				m.set(s.key, s.val, s.time)
			}

			// perform gets and ensure they are returning the expected value
			for _, g := range tst.get {
				v := m.get(g.key, g.time)
				if v != g.val {
					t.Errorf("get: %s, time: %v: expected %s but got %s", g.key, g.time, g.val, v)
					t.FailNow()
				}
			}
		})
	}
}
