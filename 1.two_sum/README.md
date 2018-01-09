# 1. Two Sum


Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

思路：

1. 两层for循环，`暴力破解`

2. 使用map数据结构。 map[数组值]数值值的索引index

if j, ok := m[target-b]; ok { // target - b为a，不断的判断a在不在构造的map中，如果不存在，将a存入map，等待之后和下一个`b`继续进行比较,最后从map中匹配到满足条件的a。则说明找到了 a + b = target 的组合，则输出对应a和b的索引。

扩展：

找值和索引对应的问题，可以借助map来额外存储，值和索引，这样也许是一个好的思路
