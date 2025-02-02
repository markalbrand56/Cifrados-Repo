package algorithms

func Substitution(s string, key string) string {
	var result string
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			result += string(key[c-'a'])
		}
		if c >= 'A' && c <= 'Z' {
			result += string(key[c-'A'])
		}
	}
	return result
}
