package algorithms

// FixedCypher Cifra un texto utilizando una clave fija.
func FixedCypher(text string, key string) string {
	var cipherText []byte

	for i := 0; i < len(text); i++ {
		// Operación XOR entre el carácter del texto y la clave
		cipherText = append(cipherText, text[i]^key[i])
	}

	return string(cipherText)
}

// DynamicCypher Cifra un texto utilizando una clave dinámica (la clave se repite si es más corta que el texto).
func DynamicCypher(text string, key string) string {
	var cipherText []byte

	keyLength := len(key)
	for i := 0; i < len(text); i++ {
		// Operación XOR entre el carácter del texto y la clave (clave se repite si es más corta)
		cipherText = append(cipherText, text[i]^key[i%keyLength])
	}

	return string(cipherText)
}

func DecryptCypher(cipherText string, key string) string {
	var text []byte

	keyLength := len(key)
	for i := 0; i < len(cipherText); i++ {
		// Operación XOR entre el carácter del texto y la clave (clave se repite si es más corta)
		text = append(text, cipherText[i]^key[i%keyLength])
	}

	return string(text)
}
