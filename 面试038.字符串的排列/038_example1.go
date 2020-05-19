package LeetCode038

// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/solution/mian-shi-ti-38-zi-fu-chuan-de-pai-lie-hui-su-fa-by/

func permutation(s string) []string {
	var res []string
	helper([]byte(s), 0, &res)
	return res
}

func helper(s []byte, start int, res *[]string) {
	if start == len(s) {
		*res = append(*res, string(s))
		return
	}
	m := make(map[byte]struct{})

	for i := start; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			continue
		}
		s[i], s[start] = s[start], s[i]
		helper(s, start+1, res)
		s[i], s[start] = s[start], s[i]
		m[s[i]] = struct{}{}
	}
}

/*
class Solution:
    def permutation(self, s: str) -> List[str]:
        c, res = list(s), []
        def dfs(x):
            if x == len(c) - 1:
                res.append(''.join(c)) # 添加排列方案
                return
            dic = set()
            for i in range(x, len(c)):
                if c[i] in dic: continue # 重复，因此剪枝
                dic.add(c[i])
                c[i], c[x] = c[x], c[i] # 交换，将 c[i] 固定在第 x 位
                dfs(x + 1) # 开启固定第 x + 1 位字符
                c[i], c[x] = c[x], c[i] # 恢复交换
        dfs(0)
				return res
*/
