# Climbing Stairs (Easy) — LeetCode #70

**Difficulty:** Easy  
**Pattern:** Dynamic programming / Fibonacci

## Data Structure

- **O(1) scalars** — two (or three) integers rolling DP state; full array optional.

## Algorithm

- `ways(n) = ways(n-1) + ways(n-2)` — reach step n from n-1 (+1 step) or n-2 (+2 steps).
- Base: `ways(1)=1`, `ways(2)=2`.
- Iterate from 3 to n updating rolling window.

**Time:** O(n)  
**Space:** O(1)

## Edge Cases

| n | Ways |
|---|------|
| 1 | 1 |
| 2 | 2 |
| 3 | 3 (1+1+1, 1+2, 2+1) |

Problem constraints usually n ≥ 1; handle n ≤ 2 explicitly.

## Go-Specific Notes

- **Integer overflow** — for very large n (not in LC constraints), use `big.Int`; n ≤ 45 fits `int`.
- **No tuple assignment** — rotate with `prev2, prev1 = prev1, curr` style manually.
- **Classic for loop** — `for i := 3; i <= n; i++` with inclusive bound; Python `range(3, n+1)`.
- Memoized recursion works but stack depth O(n); iterative is idiomatic in Go interviews.
- `int` width is platform-dependent (32 or 64 bit); LeetCode accepts 64-bit.

## Other Notes

- Generalization: k step sizes → k-term recurrence (Min Cost Climbing Stairs variant).
- Matrix exponentiation can do O(log n) for huge n (overkill here).
