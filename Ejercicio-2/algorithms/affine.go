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
	m := 26
	if gcd(a, m) != 1 {
		panic("La clave 'a' debe ser coprima con 26")
	}

	text = strings.ToUpper(text) // Convertir a mayúsculas
	var encryptedText strings.Builder

	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			P := int(char - 'A')
			C := (a*P + b) % m                     // aP + b mod m
			encryptedText.WriteByte(byte(C + 'A')) // Convertir a letra
		} else {
			encryptedText.WriteRune(char)
		}
	}
	return encryptedText.String()
}

// DecryptAffine Descifra un texto cifrado con el cifrado afín.
func DecryptAffine(text string, a int, b int) string {
	m := 26
	if gcd(a, m) != 1 {
		panic("La clave 'a' debe ser coprima con 26")
	}

	aInverse := modInverse(a, m)

	text = strings.ToUpper(text)
	var decryptedText strings.Builder

	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			C := int(char - 'A')
			P := (aInverse * ((C - b + m) % m)) % m
			decryptedText.WriteByte(byte(P + 'A'))
		} else {
			decryptedText.WriteRune(char)
		}
	}
	return decryptedText.String()
}
