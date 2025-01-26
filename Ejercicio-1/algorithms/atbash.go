package algorithms

func Atbash(s string) string {
	var result string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			result += string('z' - c + 'a')
		}
		if c >= 'A' && c <= 'Z' {
			result += string('Z' - c + 'A')
		}
	}
	return result
}
