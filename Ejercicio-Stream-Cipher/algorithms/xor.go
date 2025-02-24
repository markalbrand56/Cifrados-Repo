package algorithms

import "strings"

// XOR Función que realiza la operación XOR a una entrada binaria.
func XOR(a, b string) string {
	// Elimina los espacios en blanco que puedan tener las cadenas (separadores de bytes)
	a = strings.Replace(a, " ", "", -1)
	b = strings.Replace(b, " ", "", -1)

	var result string
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func XORBytes(data []byte, key string) []byte {
	keyBytes := []byte(key)
	keyLen := len(key)
	result := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		result[i] = data[i] ^ keyBytes[i%keyLen]
	}

	return result
}
