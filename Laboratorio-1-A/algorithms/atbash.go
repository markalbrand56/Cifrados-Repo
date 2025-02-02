package algorithms

// Atbash Cifra un texto utilizando el cifrado Atbash.
func Atbash(s string) string {
	var result string
	for _, c := range s {
		// Si el carácter es una letra minúscula, se le resta a 'z' y se suma 'a'
		if c >= 'a' && c <= 'z' {
			result += string('z' - c + 'a')
		}
		// Si el carácter es una letra mayúscula, se le resta a 'Z' y se suma 'A'
		if c >= 'A' && c <= 'Z' {
			result += string('Z' - c + 'A')
		}
	}
	return result
}
