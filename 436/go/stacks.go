package test

import "fmt"

// MultiStacker allows for stack operations on multiple stacks.
type MultiStacker interface {
	// Push pushes v onto stack n, where n is the zero based stack index
	Push(v interface{}, n uint) error
	// Pop pops a value of of stack n, where n is the zero based stack index
	Pop(n uint) (interface{}, error)
}

// NewMultiStacker returns an object that implements the MultiStacker interface with count stacks
func NewMultiStacker(count uint) (MultiStacker, error) {
	if count == 0 {
		return nil, fmt.Errorf("stack count must be nonzero")
	}

	return &stacks{
		count: count,
		elts:  make([]interface{}, count), // create with minimum length
	}, nil
}

// stacks implements MultiStacker with a single slice
// the number of stacks (count) must be known at instantiation time
type stacks struct {
	count uint
	elts  []interface{} // invariant: this should always have length >= count
}

// Push adds v to stack n, it returns an error for invalid n values
func (s *stacks) Push(v interface{}, n uint) error {
	if n > s.count-1 {
		return fmt.Errorf("no stack with index %v, valid stack indexes are: 0-%v", n, s.count-1)
	}

	var i uint
	for i = n; i < uint(len(s.elts)); i += s.count { // advance to top of stack n
		if s.elts[i] == nil { // space for v has already been allocated, but there is no value
			s.elts[i] = v
			return nil
		}
	}

	// space for i was not allocated, so allocate up to, not including index i
	for j := len(s.elts); uint(j) < i; j++ {
		s.elts = append(s.elts, nil)
	}

	s.elts = append(s.elts, v)

	return nil
}

// Pop returns the top element from stack n or an error for invalid n values
func (s *stacks) Pop(n uint) (interface{}, error) {
	if n > s.count-1 {
		return nil, fmt.Errorf("no stack with index %v, valid stack indexes are: 0-%v", n, s.count-1)
	}

	i := n
	top := s.elts[i]
	topIdx := i                                 // track the index of the "top" element in stack n
	for ; i < uint(len(s.elts)); i += s.count { // advance to top of stack n
		if s.elts[i] == nil { // previous was the top, revert i and break
			break
		}

		top = s.elts[i] // update top as current non null value
		topIdx = i
	}

	if top == nil { // there was no element to pop in stack n
		return nil, nil
	}

	s.elts[topIdx] = nil // "delete" top of stack n

	// cleanup the end of s.elts, if there are s.count nils in a row we can remove until s.elts has len = s.count
	var cur uint // cur tracks the number of nils seen in a row
	for j := len(s.elts) - 1; uint(j) > s.count-1; j-- {
		if s.elts[j] == nil {
			cur++
		} else { // found a non nil value, can't cleanup anymore
			break
		}
		if cur == s.count {
			until := uint(len(s.elts)) - s.count
			s.elts = s.elts[:until]
			cur = 0
		}
	}

	// return the top element
	return top, nil
}
