package LeetCode067

func addBinary(a string, b string) string {
	result := ""
	flag := 0
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 { // 从尾部开始,最长的字符遍历完了，就停止
		t1, t2 := 0, 0
		// 短的字符串遍历完后，继续操作，它的对应那一位其实就是0
		if i >= 0 {
			t1 = int(a[i] - '0') // 字符转为整数
		}
		if j >= 0 {
			t2 = int(b[j] - '0')
		}
		sum := t1 + t2 + flag // 计算当前位置
		switch sum {
		case 3:
			flag = 1
			result = "1" + result
		case 2:
			flag = 1
			result = "0" + result
		case 1:
			flag = 0
			result = "1" + result
		case 0:
			flag = 0
			result = "0" + result
		}
		i--
		j--
	}
	if flag == 1 { // 最终进位
		result = "1" + result
	}
	return result
}
