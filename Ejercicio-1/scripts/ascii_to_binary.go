package scripts

import (
	"fmt"
	"strconv"
	"strings"
)

// AsciiToBinary Función que convierte un texto a binario.
func AsciiToBinary(text string) string {
	var binaryResult string

	for _, char := range text {
		binaryResult += fmt.Sprintf("%08b ", char) // 08b indica que se quiere el número en binario con 8 bits.
	}

	return binaryResult
}

// BinaryToAscii Función que convierte un texto binario a texto.
func BinaryToAscii(binary string) string {
	var asciiResult strings.Builder

	// Divide la cadena binaria en bloques de 8 bits
	bits := strings.Fields(binary)

	for _, bit := range bits {
		// Verifica que el bloque tenga exactamente 8 bits
		if len(bit) != 8 {
			fmt.Println("Error: el bloque no tiene 8 bits, recuerda separar el input con espacios")
			return ""
		}

		// Convierte el bloque binario a un entero
		value, err := strconv.ParseInt(bit, 2, 64)
		if err != nil {
			fmt.Println("Error: no se pudo convertir el bloque a un entero")
			return ""
		}

		// Convierte el valor entero a un carácter ASCII
		asciiResult.WriteByte(byte(value))
	}

	return asciiResult.String()
}
