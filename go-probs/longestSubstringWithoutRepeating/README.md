# Longest Substring Without Repeating Characters (Medium) — LeetCode #3

**Difficulty:** Medium  
**Pattern:** Sliding window / hash map

## Data Structure

- **Hash map** — char → most recent index in string.
- **Window** — `[start, i]` inclusive, expanded by `i`, shrunk via `start`.

## Algorithm

1. Scan with index `i`, char `c`.
2. If `c` seen and its last index ≥ `start`, move `start` to `last[c] + 1`.
3. Update `last[c] = i`.
4. Track max window length `i - start + 1`.

**Time:** O(n)  
**Space:** O(min(n, alphabet size))

## Edge Cases

| Case | Result |
|------|--------|
| Empty string | 0 |
| All same char `"bbbbb"` | 1 |
| No repeats `"abc"` | 3 |
| Repeat outside window | Don't move start incorrectly — check `prev >= start` |
| `"abba"` | 2 (window `"ba"`) |

## Go-Specific Notes

- **`range` index on strings** — index is **byte offset**, not rune index. For ASCII inputs (LeetCode), byte index == character index. For Unicode, convert to `[]rune(s)` first or use rune-aware indexing.
- **`map[rune]int`** — keys are runes when ranging over string.
- **Condition `prev >= start`** — critical; without it, stale map entries shrink window too far.
- Alternative: `map[rune]bool` + shrink with inner loop — still valid but amortized same O(n) with two pointers.

## Other Notes

- Classic template for "longest substring with at most K distinct" (generalization).
- Set + while-loop variant is easier to explain but same complexity.
