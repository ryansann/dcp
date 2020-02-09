# Daily Coding Problem: Problem #446 [Medium]

This problem was asked by Indeed.

Given a 32-bit positive integer N, determine whether it is a power of four in faster than O(log N) time.

## Solution

```text
1   = 4^0 = 0000 0000 0000 0000 0000 0000 0000 0001
4   = 4^1 = 0000 0000 0000 0000 0000 0000 0000 0100
16  = 4^2 = 0000 0000 0000 0000 0000 0000 0001 0000
64  = 4^3 = 0000 0000 0000 0000 0000 0000 0100 0000
256 = 4^4 = 0000 0000 0000 0000 0000 0001 0000 0000
```

Notice, there is only a single bit set with the preceeding number of zeros being even when we encounter a power of 4.
