# House Robber (Medium) — LeetCode #198

**Difficulty:** Medium  
**Pattern:** 1D dynamic programming

## Data Structure

- **DP array** or two rolling variables `prev` / `curr`.
- `dp[i]` = max money robbing houses `0..i`.

## Algorithm

1. Recurrence: `dp[i] = max(dp[i-1], dp[i-2] + nums[i])`.
   - Skip house `i` → take `dp[i-1]`.
   - Rob house `i` → can't rob `i-1`, so `dp[i-2] + nums[i]`.
2. Space-optimized: track only last two values.
3. Return `curr` after iterating all houses.

**Time:** O(n)  
**Space:** O(1) with rolling vars

## Edge Cases

| Case | Notes |
|------|-------|
| Empty / single house | Return 0 or `nums[0]` |
| Two houses | `max(nums[0], nums[1])` |
| Alternating high values | DP picks optimal non-adjacent set |
| All equal | Take every other house |

## Go-Specific Notes

- **Rolling vars** — `prev, curr := 0, 0`; update with tuple-style swap in loop.
- **Slice iteration** — `for _, num := range nums`.

## Other Notes

- House Robber II (#213) adds circular street (run linear DP twice).
- House Robber III (#337) is tree DP.
