from typing import List

class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        intervals = []
        l, r = 0, 0
        diff = 0
        while r < len(nums):
            print("l, r, len(nums)", l, r, len(nums))
            print("nums[r], nums[l], diff", nums[r], nums[l], diff)
            if nums[r] == (nums[l] + diff):
                print("true")
                if r == len(nums) - 1:
                    if l == r:
                        print("single", str(nums[l]))
                        intervals.append(str(nums[l]))
                    else:
                        print("range", str(nums[l]) + "->" + str(nums[r]))
                        intervals.append(str(nums[l]) + "->" + str(nums[r]))
                r += 1
                diff += 1
            else:
                print("false")
                if r == l + 1:
                    print("single", str(nums[l]))
                    intervals.append(str(nums[l]))
                else:
                    print("range", str(nums[l]) + "->" + str(nums[r - 1]))
                    intervals.append(str(nums[l]) + "->" + str(nums[r - 1]))
                l = r
                diff = 0
        return intervals