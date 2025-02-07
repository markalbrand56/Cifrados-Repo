package algorithms

import "strings"

// Vigenere Cifra un texto utilizando el cifrado de Vigenère.
func Vigenere(text string, key string) string {
	text = strings.ToUpper(text) // Convertir a mayúsculas
	key = strings.ToUpper(key)   // Convertir a mayúsculas

	var encryptedText strings.Builder
	m := 27

	for i, char := range text {
		if (char >= 'A' && char <= 'Z') || char == 'Ñ' {
			var P, K int
			if char == 'Ñ' {
				P = 14 // Posición de la Ñ en el alfabeto español
			} else {
				P = int(char - 'A')
				if P >= 14 {
					P++ // Ajustar la posición para la Ñ
				}
			}
			K = int(key[i%len(key)] - 'A')
			if K >= 14 {
				K++ // Ajustar la posición para la Ñ
			}
			C := (P + K) % m
			if C == 14 {
				encryptedText.WriteRune('Ñ')
			} else {
				if C > 14 {
					C-- // Ajustar la posición para la Ñ
				}
				encryptedText.WriteByte(byte(C + 'A'))
			}
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
	m := 27

	for i, char := range text {
		if (char >= 'A' && char <= 'Z') || char == 'Ñ' {
			var C, K int
			if char == 'Ñ' {
				C = 14 // Posición de la Ñ en el alfabeto español
			} else {
				C = int(char - 'A')
				if C >= 14 {
					C++ // Ajustar la posición para la Ñ
				}
			}
			K = int(key[i%len(key)] - 'A')
			if K >= 14 {
				K++ // Ajustar la posición para la Ñ
			}
			P := (C - K + m) % m
			if P == 14 {
				decryptedText.WriteRune('Ñ')
			} else {
				if P > 14 {
					P-- // Ajustar la posición para la Ñ
				}
				decryptedText.WriteByte(byte(P + 'A'))
			}
		} else {
			decryptedText.WriteRune(char)
		}
	}
	return decryptedText.String()
}
