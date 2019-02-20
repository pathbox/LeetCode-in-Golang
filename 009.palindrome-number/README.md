# [9. Palindrome Number](https://leetcode.com/LeetCodes/palindrome-number/)

## 题目
Determine whether an integer is a palindrome. Do this without extra space.

## 解题思路
检查一个整数是否是回文。

先把整数转换成字符串，两个索引i，j，i从头开始走，j从尾开始走，比较索引对应的两个字符是否相同，不相同则不是回文，返回false。
