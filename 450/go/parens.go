package test

// stack is a stack of runes
type stack []rune

// push pushes a rune onto the stack
func (s *stack) push(r rune) {
	*s = append(*s, r)
}

// pop returns the top rune on the stack or 0
func (s *stack) pop() rune {
	if l := len(*s); l > 0 {
		top := (*s)[l-1]
		*s = (*s)[:l-1]
		return top
	}

	return 0
}

// rune constants
const (
	lparen   = rune('(')
	rparen   = rune(')')
	wildcard = rune('*')
)

// isBalanced returns true if s has or is capable of having balanced parenthesis
func isBalanced(s string) bool {
	return check([]rune(s), nil)
}

// check consumes rs and modifies stk
// when it encounters a wildcard, it tries to handle it as a lparen, an rparen and an empty string
// if any of these work, it returns true
func check(rs []rune, stk stack) bool {
	if len(rs) == 0 {
		if len(stk) == 0 {
			return true
		}

		return false
	}

	switch next := rs[0]; next {
	case lparen:
		stk.push(next)

		return check(rs[1:], stk)
	case rparen:
		if match := stk.pop(); match != lparen {
			return false
		}

		return check(rs[1:], stk)
	case wildcard:
		// see if making wildcard a lparen results in balanced parens
		var lpstk stack
		copy(stk, lpstk)
		lpstk.push(lparen)
		if lp := check(rs[1:], lpstk); lp == true {
			return true
		}

		// see if making wildcard a rparen results in balanced parens
		var rpstk stack
		copy(stk, rpstk)
		rpstk.pop()
		if rp := check(rs[1:], rpstk); rp == true {
			return true
		}

		return check(rs[1:], stk)
	default:
		return false // invalid rune
	}
}
