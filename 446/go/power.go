package test

import "math/bits"

// isPowerOfFour detects whether or not n is a power of 4
// it runs in constant time
func isPowerOfFour(n uint32) bool {
	if n != 0 {
		if (n & (n - 1)) == 0 { // only 1 bit is set
			// count 0 bits before the 1
			var count int
			for ; n > 1; count++ {
				n >>= 1
			}
			return count%2 == 0 // return true if even number of 0 bits before 1
		}
	}
	return false
}

// isPowerOfFourEasy outsources bitwise operations to the math/bits package
// it runs in constant time
func isPowerOfFourEasy(n uint32) bool {
	return (n != 0) && (bits.OnesCount32(n) == 1) && (bits.TrailingZeros32(n)%2 == 0)
}
