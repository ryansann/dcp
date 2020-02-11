package test

// queue is a basic queue type
type queue []interface{}

// enqueue adds an element to the back of the queue
func (q *queue) enqueue(v interface{}) {
	*q = append(*q, v)
}

// dequeue removes an element from the queue and returns it
func (q *queue) dequeue() interface{} {
	if len(*q) == 0 {
		return nil
	}

	// get dequeued element
	elt := (*q)[0]

	// remove elt from queue
	*q = (*q)[1:]

	return elt
}

// stack implements a stack with 2 queues
type stack struct {
	tmp  queue
	elts queue
}

// push pushes an element onto the stack
func (s *stack) push(v interface{}) {
	s.elts.enqueue(v) // enqueue v to the elements queue
}

// pop returns the element on top of the stack
func (s *stack) pop() interface{} {
	var top interface{}
	for next := s.elts.dequeue(); next != nil; next = s.elts.dequeue() {
		if top != nil {
			s.tmp.enqueue(top)
		}

		top = next
	}

	// swap the references for tmp and elts
	s.elts, s.tmp = s.tmp, s.elts

	return top
}
