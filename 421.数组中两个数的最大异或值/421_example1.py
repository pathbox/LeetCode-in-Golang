
# https: // leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/solution/qian-zhui-shu-by-powcai/
# 异或运算其实就是 二进制下不进位的加法
# 思路一：贪心

# 看最高位是否异或为 1,

# 比如示例中：[3, 10, 5, 25, 2, 8] -> [00011,01010, 00101, 11001, 00010, 01000]

# 从左到右最高位参看，我们发现最高位0, 0, 0, 1, 0, 0，说明最大数至少为10000

# 还有一个公式 ，如果a ^ b = c，那么c ^ a = b

# 直接看代码，理解一下吧！
#


class Solution:
    def findMaximumXOR(self, nums: List[int]) -> int:
        res = 0
        mask = 0
        for i in range(32, -1, -1):
            mask = mask | (1 << i)
            # 记录前缀
            s = set()
            for num in nums:
                s.add(num & mask)  # 得到每一位与结果，能知道该位置是1还是0
            tmp = res | (1 << i)
            for t in s:
                if tmp ^ t in s:
                    res = tmp
                    break
        return res

# 前缀树


class Solution:
    def findMaximumXOR(self, nums: List[int]) -> int:
        if not nums:
            return 0
        # 创建前缀树
        root = {}
        for num in nums:
            cur = root
            for i in range(31, -1, -1):
                cur_bit = (num >> i) & 1
                cur.setdefault(cur_bit, {})
                cur = cur[cur_bit]

        res = float("-inf")
        # 按位找最大值
        for num in nums:
            cur = root
            cur_max = 0
            for i in range(31, -1, -1):
                cur_bit = (num >> i) & 1
                if cur_bit ^ 1 in cur:
                    cur_max += (1 << i)
                    cur = cur[cur_bit ^ 1]
                else:
                    cur = cur[cur_bit]
            res = max(res, cur_max)
        return res
