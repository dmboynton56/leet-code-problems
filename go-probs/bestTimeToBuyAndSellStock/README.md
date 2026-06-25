# Best Time to Buy and Sell Stock (Easy) — LeetCode #121

**Difficulty:** Easy  
**Pattern:** Single pass / running minimum

## Data Structure

- **Scalars only** — `minPrice` and `maxProfit`; no auxiliary collections required.

## Algorithm

1. Track the lowest price seen so far (`minPrice`).
2. At each day, update `maxProfit = max(maxProfit, price - minPrice)`.
3. Buy must come before sell — single left-to-right pass enforces this.

**Time:** O(n)  
**Space:** O(1)

## Edge Cases

| Case | Result |
|------|--------|
| Monotonically decreasing prices | 0 (never sell at loss) |
| Single element | 0 |
| Empty (if allowed) | 0 |
| Best sell on last day | Handled by running max |
| Buy on day 1, sell later | Classic case |

## Go-Specific Notes

- **`math.MaxInt`** — common sentinel for "minimum so far" when starting before first element; alternatively initialize `minPrice = prices[0]`.
- **Go 1.21+** has built-in `max()`/`min()` for ordered types — older code uses manual `if`.
- **No ternary** — use `if profit > maxProfit { maxProfit = profit }`.
- Return type is `int`; profit is never negative per problem (clamp at 0).
- Empty slice: guard with `len(prices) == 0` before indexing.

## Other Notes

- Kadane-style thinking: max of `(price[i] - min(price[0..i-1]))`.
- DP variant `dp[i] = max profit ending at i` collapses to this O(1) solution.
