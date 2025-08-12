# Counting Bits

Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

[url](https://leetcode.com/problems/counting-bits/description/)

Solutions:
1. `countBits` -- cheating ssolution for go, use bits module, O(n) CPU and RAM both
2. `countBitsLow` -- solution with bits manupulation. count of 1 in actual number is count number of 1 in number with 1 right shit + odd/even option by AND operation on 1,  O(n) CPU and RAM both


tags:
`Dynamic programming` `bit manipulation`