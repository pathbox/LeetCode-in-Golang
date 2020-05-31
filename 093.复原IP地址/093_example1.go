package LeetCode093

import "strings"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}

	var result []string
	for i := 1; i < 4 && i < len(s)-2; i++ {
		for j := i + 1; j < i+4 && j < len(s)-1; j++ {
			for k := j + 1; k < j+4 && k < len(s); k++ {
				if judgeNumber(s[0:i]) && judgeNumber(s[i:j]) && judgeNumber(s[j:k]) && judgeNumber(s[k:]) {
					result = append(result, strings.Join([]string{s[0:i], s[i:j], s[j:k], s[k:]}, "."))
				}
			}
		}
	}
	return result
}

func judgeNumber(num string) bool {
	if len(num) > 1 && num[0] == '0' {
		return false
	}
	result := 0
	for i := 0; i < len(num); i++ {
		result = 10*result + int(num[i]-'0')
	}
	if result > 255 {
		return false
	}
	return true
}
