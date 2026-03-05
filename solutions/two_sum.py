from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        # TODO: Implement the solution
        map = {}
        for i in range(len(nums)):
            complement = target - nums[i]
            if complement in map:
                return [map[complement], i]
            map[nums[i]] = i
        return []
        pass
