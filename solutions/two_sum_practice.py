from typing import List

class Solution:
    def two_sum(self, nums: List[int], target: int) -> List[int]:
        complements = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in complements:
                return [complements[complement], i]
            else:
                complements[nums[i]] = i
        return []
        