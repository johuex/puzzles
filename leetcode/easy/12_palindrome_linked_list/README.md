# Palindrome Linked List

Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

[url](https://leetcode.com/problems/palindrome-linked-list/description/)

Solutions:
1. `with list` -- keep all values in array, check values from start and end to middle, O(N) -- CPU, O(N) -- RAM;
2. `with pointers` -- 1. find middle of the list; 2. reverse second part of list; 3. compare both parts of lists, O(N) -- CPU, O(1) -- RAM;


tags:
`Linked List` `Two Pointers` `Stack` `Recursion`