package LeetCode151

func reverseWords(s string) string {
	if s == "" {
		return s
	}
	res := []byte{}
	queue := []string{}

	word := []byte{}

	for i := 0; i < len(s); i++{
		if s[i] == ' ' {
			if len(word) >0 {
				queue = append(queue,string(word))
				word = []byte{}
			}else {
				word = append(word,s[i])
			}
		}
		if len(word) > 0 { // 处理最后一个单词
			queue = append(queue, string(word))
		}
		if (queue) <= 0 {
			return ""
		}
		for i := len(queue)-1; i >= 0; i--{
			res = append(res,[]byte{queue[i]...})
			res = append(res, ' ')
		}
		return string(res[:len(res)-1])
	}
}