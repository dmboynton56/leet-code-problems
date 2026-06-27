# solution returns max profit from one buy and one sell.
# Go: math.MaxInt for initial min; Python uses float('inf') or nums[0].
def solution(prices: list[int]) -> int:
    if not prices:
        return 0

    min_price = prices[0]  # Go uses math.MaxInt as sentinel; Python starts at first price.
    max_profit = 0

    for price in prices:
        # Go: separate if price < minPrice and else-if profit > maxProfit branches.
        # Python: built-in min/max in one loop body (same outcome, more compact).
        min_price = min(min_price, price)
        max_profit = max(max_profit, price - min_price)

    return max_profit


if __name__ == "__main__":
    print(solution([7, 1, 5, 3, 6, 4]))  # 5
    print(solution([7, 6, 4, 3, 1]))     # 0
