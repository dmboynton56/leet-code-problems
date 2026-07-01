# Longest Consecutive Sequence (Medium) — LeetCode #128

**Difficulty:** Medium  
**Pattern:** Hash set / smart enumeration

## Data Structure

- **Hash set** of all numbers for O(1) lookup.
- Only start counting from **sequence starts** (no `num-1` in set).

## Algorithm

1. Insert all nums into a set.
2. For each `num` in set:
   - If `num-1` not in set, `num` is start of a sequence.
   - Walk `num, num+1, num+2, ...` while in set; track length.
   - Update global max length.
3. Return max length.

**Time:** O(n) — each element visited at most twice  
**Space:** O(n)

## Edge Cases

| Case | Notes |
|------|-------|
| Empty array | 0 |
| Single element | 1 |
| Duplicates in input | Set dedupes automatically |
| Negative numbers | Works unchanged |
| Unsorted input | No sort needed — key advantage |

## Go-Specific Notes

- **Set idiom** — `map[int]struct{}` or `map[int]bool`; empty struct saves memory.
- **Range over set** — `for num := range seen`.
- Sort + scan is O(n log n); set approach hits O(n) requirement.

## Other Notes

- Brute force sort + linear scan is O(n log n), usually acceptable but not optimal.
- Union-Find also O(n α(n)) but set approach is simpler.
