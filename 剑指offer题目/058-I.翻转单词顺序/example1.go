package Offer058

import "strings"

/*
待转换字符串str = "yang zhi ju"
1. 先将每个单词旋转一遍：
"yang" -> "gnay"
"zhi"  -> "ihz"
"ju"   -> "uj"
最后得到字符串str2 = "gnay ihz uj"

2. 将str2旋转一遍得到str3 = "ju zhi yang"
*/

func reverseWords(s string) string {
	var reserseSeg []string
	//分割字符串
	seg := strings.Fields(s)
	//反转字符串数组
	for i := len(seg) - 1; i >= 0; i-- {
		reserseSeg = append(reserseSeg, seg[i])
	}
	//将字符串数组里的元素拼接成一个字符串并返回
	return strings.Join(reserseSeg, " ")
}

/*
class Solution:
    def reverseWords(self, s: str) -> str:
        """按空格分割成单词列表，1.末位单词+空格 2.最后一个单词末位不加空格"""
        if s == [] : return ""
        ls = s.split()  # ['the', 'sky', 'is', 'blue']
        if ls == []: return ""
        res = ""
        for i in range(len(ls)-1): # 少一个，处理末位不加空格
            res += ls[len(ls) - 1 - i] + " "  # 最后一个单词 + 空格
        res += ls[0] # 末位不加空格
        return res
*/
