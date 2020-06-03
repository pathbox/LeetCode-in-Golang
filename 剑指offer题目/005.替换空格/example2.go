package Offer005

func replaceSpace(s string) string {
	var b []rune
	for _, v := range s {
		if v == 32 {
			b = append(b, 37, 50, 48)
		} else {
			b = append(b, v)
		}
	}
	return string(b)
}
