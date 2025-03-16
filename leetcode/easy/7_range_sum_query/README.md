# Range Sum Query - Immutable

Given an integer array nums, handle multiple queries of the following type:

Calculate the sum of the elements of nums between indices left and right inclusive where left <= right.
Implement the NumArray class:

NumArray(int[] nums) Initializes the object with the integer array nums.
int sumRange(int left, int right) Returns the sum of the elements of nums between indices left and right inclusive (i.e. nums[left] + nums[left + 1] + ... + nums[right]).

[url](https://leetcode.com/problems/range-sum-query-immutable/description/)

Solutions:
1. `stupid solution` --> CPU: O(n), RAM: O(n), we calculate each range with loop.
2. `best solution` --> CPU: O(n), RAM: O(n), we calculate sum of range ony once, later for each range we calculate only once.

__P.S__
Both solution is O(n), but in second solution it's clear O(n) but in first solution it'll be M*O(n) where M -- count of testCases

tags:
`array` `design` `prefix sum`