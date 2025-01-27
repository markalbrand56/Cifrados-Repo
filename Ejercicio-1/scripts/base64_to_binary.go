package scripts

import "strings"

// Base64ToBinary Transforma un caracter en base64 a binario.
func binary(n rune) string {
	var result string
	for i := 5; i >= 0; i-- {
		bit := n & (1 << uint(i))
		if bit > 0 {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
}

func Base64ToBinary(s string) string {
	var result string
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			result += binary(c - 'A')
		}
		if c >= 'a' && c <= 'z' {
			result += binary(c - 'a')
		}
		if c >= '0' && c <= '9' {
			result += binary(c - '0')
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

func BinaryToBase64(s string) string {
	s = strings.Replace(s, " ", "", -1) // Elimina los espacios en blanco (separadores de bytes)

	var result string
	for i := 0; i < len(s); i += 6 {
		var n rune
		for j := 0; j < 6; j++ {
			n <<= 1
			if i+j < len(s) && s[i+j] == '1' {
				n |= 1
			}
		}
		if n < 26 {
			result += string(n + 'A')
		}
		if n >= 26 && n < 52 {
			result += string(n - 26 + 'a')
		}
		if n >= 52 && n < 62 {
			result += string(n - 52 + '0')
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
