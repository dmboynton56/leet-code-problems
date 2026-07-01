# Container With Most Water (Medium) — LeetCode #11

**Difficulty:** Medium  
**Pattern:** Two pointers / greedy

## Data Structure

- **Two indices** `left` and `right` bounding the current container.
- **Running max** for the best area seen.

## Algorithm

1. Initialize `left = 0`, `right = n-1`, `maxArea = 0`.
2. While `left < right`:
   - Height of container = `min(height[left], height[right])`.
   - Width = `right - left`.
   - Update `maxArea`.
   - Move the pointer at the **shorter** line inward (the taller side cannot improve area while width shrinks).
3. Return `maxArea`.

**Time:** O(n)  
**Space:** O(1)

## Edge Cases

| Case | Notes |
|------|-------|
| Two elements | Only one container to evaluate |
| All equal heights | Any pair gives same area; greedy still correct |
| Monotonic increasing/decreasing | Shorter side always moves |
| Very skewed heights | Greedy proof: discarding shorter side is safe |

## Go-Specific Notes

- **`min(a, b)`** — Go 1.21+ has built-in `min`; older code uses inline `if`.
- **Slice bounds** — `height []int`; no need to copy input.
- Integer overflow unlikely given constraints; area fits in `int`.

## Other Notes

- Brute force O(n²) pairs works but is suboptimal.
- Different from Trapping Rain Water (#42): here you pick two walls; there you sum water between many walls.
