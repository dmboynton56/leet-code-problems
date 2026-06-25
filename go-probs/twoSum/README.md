# Two Sum (Easy) — LeetCode #1

**Difficulty:** Easy  
**Pattern:** Hash map / one-pass lookup

## Data Structure

- **Hash map** (`map[int]int` in Go, `dict[int, int]` in Python): maps each value seen so far to its index.

## Algorithm

1. Iterate the array once.
2. For each element `num`, compute `complement = target - num`.
3. If `complement` is already in the map, return `[map[complement], current_index]`.
4. Otherwise store `num -> index` and continue.
5. Return empty if no pair exists (problem guarantees exactly one solution, so this is defensive).

**Time:** O(n)  
**Space:** O(n)

## Edge Cases

| Case | Notes |
|------|-------|
| Duplicate values | e.g. `[3, 3], target=6` — store index on first pass, hit on second |
| Negative numbers / zero | Works unchanged; map keys are ints |
| Same element twice | Problem says you may not use the same element twice; storing after check avoids this |
| Two elements total | Trivial base case |

## Go-Specific Notes

- **`make(map[int]int)`** — writing to a nil map panics; always initialize.
- **Comma-ok idiom** — `if j, ok := seen[complement]; ok` is idiomatic; Python uses `in`.
- **Slice return** — `[]int{j, i}` is a composite literal; order matches LeetCode's expected `[i, j]`.
- **`range`** — gives index and value; equivalent to Python's `enumerate`.
- No built-in `defaultdict`; plain map is fine here.
- LeetCode Go stubs often use `func twoSum(nums []int, target int) []int` — rename as needed for submission.

## Other Notes

- Brute force O(n²) double loop works but is interview-suboptimal.
- Sort + two pointers works only if you track original indices (different problem variant).
