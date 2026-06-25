# Product of Array Except Self (Medium) — LeetCode #238

**Difficulty:** Medium  
**Pattern:** Prefix/suffix product / two passes

## Data Structure

- **Output array** — same length as input; holds final products.
- **Scalar `suffix`** — running product from the right (prefix folded into first pass).

## Algorithm

1. **Forward:** `result[i]` = product of all elements left of `i`. Set `result[0]=1`, then `result[i] = result[i-1] * nums[i-1]`.
2. **Backward:** track `suffix` = product of elements right of `i`; `result[i] *= suffix`; update `suffix *= nums[i]`.

**Time:** O(n)  
**Space:** O(1) extra (output array doesn't count per problem)

## Edge Cases

| Case | Notes |
|------|-------|
| Contains zero | One index gets non-zero product, others get 0 |
| Multiple zeros | All products 0 |
| Negative numbers | Sign flips handled by multiplication |
| Two elements | `[a,b]` → `[b,a]` |
| Single element | Product of empty set = 1 (problem usually n ≥ 2) |

## Go-Specific Notes

- **`make([]int, n)`** — zero-initialized; set `result[0] = 1` explicitly before loop.
- **Reverse loop** — `for i := n - 1; i >= 0; i--` ; Python `range(n-1, -1, -1)`.
- **No division** — constraint in problem; also avoids divide-by-zero with zeros in array.
- Watch **overflow** — product of many large ints; LeetCode accepts 32-bit int results typically.
- Alternative: separate prefix and suffix arrays — O(n) space, easier to read.

## Other Notes

- Follow-up "O(1) extra space" means excluding output — this solution qualifies.
- Log-sum trick with division if zeros absent — not allowed here.
