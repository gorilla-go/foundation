package strings

func UcFirst(s string) string {
	var upperStr string
	vv := []rune(s)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return s
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
