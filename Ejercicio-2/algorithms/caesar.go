package algorithms

func Caesar(s string, shift int) string {
	shift = shift % 26
	var result string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c += rune(shift)
			if c > 'z' {
				c -= 26
			}
		}
		if c >= 'A' && c <= 'Z' {
			c += rune(shift)
			if c > 'Z' {
				c -= 26
			}
		}
		result += string(c)
	}
	return result
}

func DecryptCaesar(s string, shift int) string {
	shift = shift % 26
	var result string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c -= rune(shift)
			if c < 'a' {
				c += 26
			}
		}
		if c >= 'A' && c <= 'Z' {
			c -= rune(shift)
			if c < 'A' {
				c += 26
			}
		}
		result += string(c)
	}
	return result
}
