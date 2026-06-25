# 3Sum (Medium) — LeetCode #15

**Difficulty:** Medium  
**Pattern:** Sort + two pointers

## Data Structure

- **Sorted array** — enables two-pointer search and duplicate skipping.
- **Result** — slice of triplets `[][]int`.

## Algorithm

1. Sort `nums`.
2. For each index `i` (anchor):
   - Skip if `nums[i] == nums[i-1]` (duplicate anchor).
   - Two pointers `left = i+1`, `right = n-1`.
   - If sum == 0: record triplet, advance both, skip duplicate left/right values.
   - If sum < 0: `left++`; else `right--`.
3. Stop anchor when `i > n-3`.

**Time:** O(n²)  
**Space:** O(1) extra excluding output and sort stack O(n log n)

## Edge Cases

| Case | Notes |
|------|-------|
| Fewer than 3 elements | Return empty |
| All zeros `[0,0,0,0]` | One triplet `[0,0,0]` |
| No solution | `[]` |
| Many duplicates | Duplicate-skipping loops essential |
| Large negatives/positives | Sort handles ordering |

## Go-Specific Notes

- **`sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })`** — in-place sort with custom less; import `"sort"`.
- **`append(result, []int{...})`** — appending slice to `[][]int`; preallocate with `make` if optimizing.
- **Duplicate skip** — after moving pointers, inner `for` loops skip equal neighbors; easy to miss in Go port from Python.
- Return empty as `[][]int{}` or `nil` — both serialize to `[]` in JSON.
- Avoid using `map[string]bool` for dedup if sort+skip done correctly — cleaner O(1) extra.

## Other Notes

- 3Sum Closest / 4Sum are natural extensions (fix more anchors, two pointers on remainder).
- Hash map approach for fixed `i` reduces to Two Sum II on sorted tail — same complexity.
