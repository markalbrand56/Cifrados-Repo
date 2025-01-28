package scripts

import "strings"

// Base64ToBinary Transforma un caracter en base64 a binario.
func binary(n rune) string {
	var result string

	// Iterar por los 6 bits del caracter
	for i := 5; i >= 0; i-- {
		bit := n & (1 << uint(i)) // Extraer el bit i

		if bit > 0 {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

// Base64ToBinary Transforma un texto en base64 a binario.
func Base64ToBinary(s string) string {
	var result string

	for _, c := range s {
		if c == '=' {
			// El caracter '=' es un separador de bytes, por lo que se ignora
			continue
		}

		if c >= 'A' && c <= 'Z' {
			result += binary(c - 'A') // 0-25
		}
		if c >= 'a' && c <= 'z' {
			result += binary(c - 'a') // 26-51
		}
		if c >= '0' && c <= '9' {
			result += binary(c - '0') // 52-61
		}
		if c == '+' {
			result += binary(62)
		}
		if c == '/' {
			result += binary(63)
		}
	}
	return result
}

// BinaryToBase64 Transforma un texto binario a base64.
func BinaryToBase64(s string) string {
	s = strings.Replace(s, " ", "", -1) // Elimina los espacios en blanco (separadores de bytes)

	var result string

	// Iterar por bloques de 6 bits
	for i := 0; i < len(s); i += 6 {
		var n rune // Número de 6 bits

		// Construir el número de 6 bits
		for j := 0; j < 6; j++ {
			n <<= 1

			// Verificar si el bit j está activo
			if i+j < len(s) && s[i+j] == '1' {
				n |= 1 // Encender el bit menos significativo
			}
		}

		if n < 26 {
			result += string(n + 'A') // A-Z
		}
		if n >= 26 && n < 52 {
			result += string(n - 26 + 'a') // a-z
		}
		if n >= 52 && n < 62 {
			result += string(n - 52 + '0') // 0-9
		}
		if n == 62 {
			result += "+"
		}
		if n == 63 {
			result += "/"
		}
	}
	return result
}
