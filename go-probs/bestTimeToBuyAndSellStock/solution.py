def solution(prices: list[int]) -> int:
    """
    Track running minimum buy price and max profit so far.
    Python: max() built-in; Go uses explicit if or math.Max (Go 1.21+ has max() builtin).
    """
    if not prices:
        return 0

    min_price = prices[0]
    max_profit = 0

    for price in prices:
        min_price = min(min_price, price)
        max_profit = max(max_profit, price - min_price)

    return max_profit


if __name__ == "__main__":
    print(solution([7, 1, 5, 3, 6, 4]))  # 5
    print(solution([7, 6, 4, 3, 1]))     # 0
