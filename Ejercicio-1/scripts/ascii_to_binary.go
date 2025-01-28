package scripts

import (
	"fmt"
	"strconv"
	"strings"
)

// AsciiToBinary Función que convierte un texto a binario.
func AsciiToBinary(text string) string {
	var binaryResult string

	// char es un rune, que es un alias para int32 lo que significa que se está tomando su valor ASCII.
	for _, char := range text {
		binaryResult += fmt.Sprintf("%08b ", char) // 08b indica que se quiere el número en binario con 8 bits.
	}

	return binaryResult
}

// BinaryToAscii Función que convierte un texto binario a texto.
func BinaryToAscii(binary string) string {
	binary = strings.ReplaceAll(binary, " ", "")

	var asciiResult strings.Builder

	// Verifica que la longitud de la cadena binaria sea múltiplo de 8
	if len(binary)%8 != 0 {
		fmt.Println("Error: la longitud de la cadena binaria no es múltiplo de 8")
		return ""
	}

	// Itera sobre la cadena binaria en bloques de 8 bits
	for i := 0; i < len(binary); i += 8 {
		bit := binary[i : i+8]

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
