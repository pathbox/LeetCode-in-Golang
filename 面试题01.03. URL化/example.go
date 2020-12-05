package Offer0103

func replaceSpaces(S string, length int) string {
	var arr []byte
	for i := 0; i < length; i++ {
		if i >= len(S) || S[i] == 32 {
			arr = append(arr, '%', '2', '0')
		} else {
			arr = append(arr, S[i])
		}
	}
	return string(arr)
	// return strings.Replace(S[:length], " ", "%20", -1)
}
