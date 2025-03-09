# Find all disappeared numbers in array

Given an array nums of n integers where nums[i] is in the range [1, n], return an array of all the integers in the range [1, n] that do not appear in nums.

[url](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/)

Solutions:
1. `findDisappearedNumbers` - CPU: O(n), RAM: O(n), two loops + hashmap
2. `findDisappearedNumbers` - CPU: O(n), RAM: O(1), two loops only. Solution: make nums[num-1] as negative and than search index of positive numbers.

tags:
`array` `hash table`