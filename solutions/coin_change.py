from typing import List

"""
Problem: Coin Change
You are given an integer array coins representing coins of different denominations and an integer 
amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. If that amount of money 
cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1

Example 3:
Input: coins = [1], amount = 0
Output: 0
"""

class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        # dp[i] will store the minimum number of coins needed for amount i
        # Initialize with amount + 1 (an impossible value) as a placeholder
        dp = [amount + 1] * (amount + 1)
        
        # Base case: 0 coins are needed to make an amount of 0
        dp[0] = 0
        
        # Iterate through every sub-amount from 1 up to our target amount
        for a in range(1, amount + 1):
            # Try every available coin denomination for the current sub-amount
            for coin in coins:
                # If the coin is smaller than or equal to the current sub-amount
                if a - coin >= 0:
                    # Update dp[a] if using this coin plus the optimal solution 
                    # for the remaining amount (a - coin) is better (fewer coins)
                    dp[a] = min(dp[a], 1 + dp[a - coin])
        
        # If dp[amount] is still the placeholder, the amount cannot be made
        if dp[amount] == amount + 1:
            return -1
            
        # Return the minimum coins calculated for the target amount
        return dp[amount]

    pass

