package test

import (
	"reflect"
	"testing"
)

type itineraryT struct {
	flights []flight
	start   string
	expect  []string
}

func TestFindItinerary(t *testing.T) {
	tests := []itineraryT{
		{ // single possible itinerary
			flights: []flight{
				{"SFO", "HKO"},
				{"YYZ", "SFO"},
				{"YUL", "YYZ"},
				{"HKO", "ORD"},
			},
			start:  "YUL",
			expect: []string{"YUL", "YYZ", "SFO", "HKO", "ORD"},
		},
		{ // no such itinerary exists
			flights: []flight{
				{"SFO", "COM"},
				{"COM", "YYZ"},
			},
			start:  "COM",
			expect: nil,
		},
		{ // test lexicographical requirement (multiple possible orderings)
			flights: []flight{
				{"A", "B"},
				{"A", "C"},
				{"B", "C"},
				{"C", "A"},
			},
			start:  "A",
			expect: []string{"A", "B", "C", "A", "C"}, // not []string{"A", "C", "A", "B", "C"}
		},
		{
			flights: []flight{
				{"A", "B"},
				{"B", "A"},
				{"A", "B"},
				{"B", "A"},
				{"A", "B"},
				{"B", "A"},
				{"A", "B"},
			},
			start:  "A",
			expect: []string{"A", "B", "A", "B", "A", "B", "A", "B"},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			t.Logf("flights: %v, start: %v", test.flights, test.start)

			itinerary := findItinerary(test.flights, test.start)
			if !reflect.DeepEqual(itinerary, test.expect) {
				t.Fatalf("expected: %v, got: %v", test.expect, itinerary)
			}

			t.Logf("got: %v", itinerary)
		})
	}
}
