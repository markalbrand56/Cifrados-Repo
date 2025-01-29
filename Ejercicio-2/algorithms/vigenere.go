package algorithms

import "strings"

// Vigenere Cifra un texto utilizando el cifrado de Vigenère.
func Vigenere(text string, key string) string {
	text = strings.ToUpper(text) // Convertir a mayúsculas
	key = strings.ToUpper(key)   // Convertir a mayúsculas

	var encryptedText strings.Builder

	for i, char := range text {
		if char >= 'A' && char <= 'Z' {
			P := int(char - 'A')
			K := int(key[i%len(key)] - 'A')
			C := (P + K) % 26
			encryptedText.WriteByte(byte(C + 'A'))
		} else {
			encryptedText.WriteRune(char)
		}
	}
	return encryptedText.String()
}

// DecryptVigenere Descifra un texto cifrado con el cifrado de Vigenère.
func DecryptVigenere(text string, key string) string {
	text = strings.ToUpper(text)
	key = strings.ToUpper(key)
	var decryptedText strings.Builder

	for i, char := range text {
		if char >= 'A' && char <= 'Z' {
			C := int(char - 'A')
			K := int(key[i%len(key)] - 'A')
			P := (C - K + 26) % 26
			decryptedText.WriteByte(byte(P + 'A'))
		} else {
			decryptedText.WriteRune(char)
		}
	}
	return decryptedText.String()
}
