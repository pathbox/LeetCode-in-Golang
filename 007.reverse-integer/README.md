# [7. Reverse Integer](https://leetcode.com/LeetCodes/reverse-integer/)

## 题目
Reverse digits of an integer.
```
Example1: x = 123, return 321
Example2: x = -123, return -321
```
## 解题思路

对于一个整数x

- 取末尾的数的方法

  r = x % 10 得到末尾的数

- 去除末尾的数的方法

  r = x / 10 得到去除末尾后的数