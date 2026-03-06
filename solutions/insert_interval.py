from typing import List

"""
Problem: Insert Interval
You are given an array of non-overlapping intervals where intervals[i] = [starti, endi] 
representing the start and the end of the ith interval and intervals is sorted 
in ascending order by starti. You are also given an interval newInterval = [start, end] 
that represents the start and end of another interval.

Insert newInterval into intervals such that intervals is still sorted in 
ascending order by starti and intervals still does not have any overlapping 
intervals (merge overlapping intervals if necessary).

Return intervals after the insertion.

Example 1:
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]

Example 2:
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
"""

class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        res = []
        i = 0
        n = len(intervals)
        
        # 1. Add all intervals ending before newInterval starts
        while i < n and intervals[i][1] < newInterval[0]:
            res.append(intervals[i])
            i += 1
            
        # 2. Merge overlapping intervals
        while i < n and intervals[i][0] <= newInterval[1]:
            newInterval[0] = min(newInterval[0], intervals[i][0])
            newInterval[1] = max(newInterval[1], intervals[i][1])
            i += 1
        res.append(newInterval)
        
        # 3. Add remaining intervals
        while i < n:
            res.append(intervals[i])
            i += 1
            
        return res