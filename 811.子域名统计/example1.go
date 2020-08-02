package LeetCode811

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	count := make(map[string]int, 0)
	for _, value := range cpdomains {
		splits := strings.Split(value, " ")
		num, _ := strconv.Atoi(splits[0])
		for {
			count[splits[1]] += num // splits[1] 是域名 作为key
			//找当前域名第一个点的位置
			dotIndex := strings.Index(splits[1], ".")
			//找不到点说明这一条记录完顶级域名了
			if dotIndex < 0 {
				break
			}
			//裁剪点前面的部分得到更高一级域名
			splits[1] = splits[1][dotIndex+1:]
		}
	}
	ans := make([]string, 0, len(count))
	for key, value := range count {
		ans = append(ans, fmt.Sprintf("%d %s", value, key))
	}
	return ans
}
