# Search in Rotated Sorted Array (Medium) — LeetCode #33

**Difficulty:** Medium  
**Pattern:** Binary search on rotated array

## Data Structure

- **Sorted array** rotated at unknown pivot — still has two sorted halves.
- **Two pointers** `left`, `right` for binary search.

## Algorithm

1. Standard binary search loop while `left <= right`.
2. Compute `mid`. If `nums[mid] == target`, return `mid`.
3. Determine which half is sorted:
   - If `nums[left] <= nums[mid]`: left half sorted.
     - If `target` in `[nums[left], nums[mid])`, search left; else right.
   - Else: right half sorted.
     - If `target` in `(nums[mid], nums[right]]`, search right; else left.
4. Return -1 if not found.

**Time:** O(log n)  
**Space:** O(1)

## Edge Cases

| Case | Notes |
|------|-------|
| Not rotated (pivot 0) | Degrades to normal binary search |
| Single element | Direct compare |
| Target at pivot | Handled by mid check |
| Duplicates | This problem has unique elements; duplicates need extra logic (#81) |

## Go-Specific Notes

- **Half-open intervals** — careful with `<=` vs `<` when checking target in range.
- **Integer mid** — `mid := left + (right-left)/2` avoids overflow (Go ints don't overflow same way but idiom is standard).

## Other Notes

- Find Minimum in Rotated Sorted Array (#153) is the companion problem.
- Key insight: one half is always sorted — use that to decide which side to discard.
