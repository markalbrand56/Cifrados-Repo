package algorithms

import (
	"strings"
)

// gcd Calcula el máximo común divisor de dos números.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// modInverse Calcula el inverso multiplicativo de un número en módulo m.
func modInverse(a, m int) int {
	for i := 0; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	panic("No existe inverso multiplicativo para a en módulo m")
}

// Affine Cifra un texto utilizando el cifrado afín.
func Affine(text string, a int, b int) string {
	m := 27
	if gcd(a, m) != 1 {
		panic("La clave 'a' debe ser coprima con 27")
	}

	text = strings.ToUpper(text) // Convertir a mayúsculas
	var encryptedText strings.Builder

	for _, char := range text {
		if (char >= 'A' && char <= 'Z') || char == 'Ñ' {
			var P int
			if char == 'Ñ' {
				P = 14 // Posición de la Ñ en el alfabeto español
			} else {
				P = int(char - 'A')
				if P >= 14 {
					P++ // Ajustar la posición para la Ñ
				}
			}
			C := (a*P + b) % m
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

// DecryptAffine Descifra un texto cifrado con el cifrado afín.
func DecryptAffine(text string, a int, b int) string {
	m := 27
	if gcd(a, m) != 1 {
		panic("La clave 'a' debe ser coprima con 27")
	}

	aInverse := modInverse(a, m)

	text = strings.ToUpper(text)
	var decryptedText strings.Builder

	for _, char := range text {
		if (char >= 'A' && char <= 'Z') || char == 'Ñ' {
			var C int
			if char == 'Ñ' {
				C = 14 // Posición de la Ñ en el alfabeto español
			} else {
				C = int(char - 'A')
				if C >= 14 {
					C++ // Ajustar la posición para la Ñ
				}
			}
			P := (aInverse * ((C - b + m) % m)) % m
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
