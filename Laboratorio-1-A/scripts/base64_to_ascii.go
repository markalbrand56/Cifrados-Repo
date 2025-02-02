package scripts

import (
	"strings"
)

var base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Base64ToAscii Transforma un texto en base64 a texto.
func Base64ToAscii(encoded string) string {
	var decoded []byte
	var buffer int
	var bitsCount int

	// Elimina los caracteres de relleno '='
	encoded = strings.TrimRight(encoded, "=")

	for _, char := range encoded {
		// Encuentra el índice del carácter en la tabla Base64
		index := strings.IndexRune(base64Table, char)
		if index == -1 {
			return ""
		}

		// Acumula los bits
		buffer = (buffer << 6) | index
		bitsCount += 6

		// Extrae los bytes completos (8 bits cada uno)
		for bitsCount >= 8 {
			bitsCount -= 8
			decoded = append(decoded, byte((buffer>>bitsCount)&0xFF))
		}
	}

	return string(decoded)
}

// AsciiToBase64 Transforma un texto a base64.
func AsciiToBase64(text string) string {
	var base64Result strings.Builder
	var buffer int
	var bitsCount int

	for _, char := range text {
		// Agrega los bits del carácter ASCII al buffer
		buffer = (buffer << 8) | int(char)
		bitsCount += 8

		// Mientras tengamos al menos 6 bits, generamos un carácter Base64
		for bitsCount >= 6 {
			bitsCount -= 6
			index := (buffer >> bitsCount) & 0x3F // Extrae los 6 bits más significativos
			base64Result.WriteByte(base64Table[index])
		}
	}

	// Procesar los bits restantes
	if bitsCount > 0 {
		buffer <<= 6 - bitsCount // Desplaza los bits restantes a la izquierda
		base64Result.WriteByte(base64Table[buffer&0x3F])
	}

	// Agregar caracteres de relleno '=' si es necesario
	for base64Result.Len()%4 != 0 {
		base64Result.WriteByte('=')
	}

	return base64Result.String()
}
