package LeetCode139

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j, v := range wordDict {
			l := len(v) // 每个单词的长度
			if dp[i] == true && (i+1 <= n && s[i:i+l] == wordDict[j]) {
				dp[i+1] = true
			}
		}
	}
	return dp[n]
}

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true // 每个单词存到map中
	}
	sLength := len(s)
	dp := make([]bool, sLength+1)
	//dp[i]表示下标0...i-1所截取的s串能否做"单词拆分"，true表示能，false表示不能
	dp[0] = true
	//结合dp[i]的意义，i=1表示检查下标0...0所截取的s串(即s[0])；i=len(s)表示检查下标0...len(s)-1所截取的s串(即整个s串)
	//从s[0]一直遍历到s[0:len(s)-1]
	for i := 1; i <= sLength; i++ {
		for j := 0; j < i; j++ {
			//对于下标0...i-1所截取的s串，将其分为0...j-1和j...i-1两个部分。
			//dp[j]表示下标0...j-1所截取的s串能否做"单词拆分"，true表示能，false表示不能
			//若dp[j]=true,并且s[j:i]是wordDict里的一个单词(表现为m[s[j:i]]=true),则dp[i]=true
			if dp[j] && m[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[sLength]
}
