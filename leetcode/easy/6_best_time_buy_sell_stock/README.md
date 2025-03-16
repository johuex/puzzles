# Best Time to Buy and Sell Stock

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

[url](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/)

Solutions:
1. `maxProfit` --> CPU: O(nlogn), RAM: O(n), store all diffs between min and current, sort and select [0].
1. `maxProfitOpt` --> CPU: O(n), RAM: O(1), calculate profit between minimum and price and store only it.

tags:
`array` `dynamic programming`