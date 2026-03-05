from typing import List

"""
Problem: Group Anagrams
Given an array of strings strs, group the anagrams together. You can return the answer in any order.
An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
typically using all the original letters exactly once.

Example 1:
Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:
Input: strs = [""]
Output: [[""]]

Example 3:
Input: strs = ["a"]
Output: [["a"]]
"""
from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # first instinct is map that holds count of each letter in each word, so as looping through the strings, reset map and check all other strings
        # for strings proven anagrams of each other, add to anagramMap, so don't have to loop through all strings for those again
        # charMap = {}
        # anagramMap = {}
        # for s in strs:
        #     for ch in s:
        #         if charMap[ch] > 0: # curr character exists in word we are checking
        #             charMap[ch] -= 1
        #         else: # doesn't have current letter, isn't an anagram
        #             continue # or break ?
        #     anagramMap[s] += s # add str we are checking as anagram of current string
        # return anagramMap # how to then group anagrams into list ?  
        
        anagram_map = defaultdict(list)
        
        for s in strs:
            # 1. Create the signature (sorted string)
            # 2. Use it as the key to group the original string
            key = "".join(sorted(s))
            anagram_map[key].append(s)
            
        # Return just the lists of groups
        return list(anagram_map.values())
