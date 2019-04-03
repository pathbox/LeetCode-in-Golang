# [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)

## 题目
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

## 解题思路

```
func climbStairs_recursion(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {  // 返回条件
		return 1
	}

	r := climbStairs(n-1) + climbStairs(n-2) // 爬到n处，可能是从n-1处过来的，也可能是从n-2处过来的
	return r
}
```