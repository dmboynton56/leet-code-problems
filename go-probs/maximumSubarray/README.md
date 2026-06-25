# Maximum Subarray (Easy) — LeetCode #53

**Difficulty:** Easy  
**Pattern:** Kadane's algorithm / DP

## Data Structure

- **Scalars** — `currentSum` (best sum ending here), `maxSum` (global best).

## Algorithm (Kadane)

For each element:
1. `currentSum = max(nums[i], currentSum + nums[i])` — extend or restart.
2. `maxSum = max(maxSum, currentSum)`.

Equivalent: if `currentSum < 0`, reset before adding next (when all negative handled by init).

**Time:** O(n)  
**Space:** O(1)

## Edge Cases

| Case | Notes |
|------|-------|
| All negative | Answer is the least negative single element |
| Single element | That element |
| Entire array is best | maxSum updates throughout |
| Zeros | Fine; subarray can include zeros |

## Go-Specific Notes

- **Initialize from `nums[0]`** — problem guarantees non-empty array; safe to index `[0]`.
- **Go 1.21+ `max()`** — can write `currentSum = max(nums[i], currentSum+nums[i])`.
- Watch **integer overflow** if summing huge values (not typical in LC).
- Divide-and-conquer O(n log n) exists but Kadane is the expected answer.

## Other Notes

- Follow-up: return actual subarray indices — track start/end when max updates.
- Related: Maximum Product Subarray (track min and max because negatives flip sign).
