# Single Number

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

[url](https://leetcode.com/problems/single-number/description/)

Solutions:
1. `singleNumber` --> CPU: O(nlogn), RAM: O(1), Sort + one loop
1. `singleNumberXOR` --> CPU: O(N), RAM: O(1), The XOR ( ^ ) is an logical operator that will return 1 when the bits are different and 0 elsewhere, only one loop.

tags:
`array` `math` `bit manipulation`