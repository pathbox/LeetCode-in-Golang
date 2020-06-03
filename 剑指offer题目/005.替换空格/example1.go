package Offer005

func replaceSpace(s string) string {
	r := ""
	l := len(s)
	for i := 0; i <= l-1; i++ {
		tmp := string(s[i])
		if tmp == " " {
			tmp = "%20"
		}
		r = r + tmp
	}
	return r
}
