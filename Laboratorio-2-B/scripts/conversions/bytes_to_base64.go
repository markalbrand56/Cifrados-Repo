package conversions

import (
	"errors"
	"strings"
)

// BytesToBase64 Función que convierte un arreglo de bytes a una cadena en Base64.
func BytesToBase64(data []byte) string {
	var encoded string
	padding := 0

	// Procesar cada bloque de 3 bytes
	for i := 0; i < len(data); i += 3 {
		var chunk uint32

		// Cargar hasta 3 bytes en un entero de 24 bits
		for j := 0; j < 3; j++ {
			chunk <<= 8
			if i+j < len(data) {
				chunk |= uint32(data[i+j])
			} else {
				padding++
			}
		}

		// Extraer 4 bloques de 6 bits y mapear a la tabla Base64
		for j := 3; j >= 0; j-- {
			if i+(3-j) < len(data)+padding {
				encoded += string(base64Table[(chunk>>(j*6))&0x3F])
			} else {
				encoded += "="
			}
		}
	}

	return encoded
}

// base64Map Mapeo de caracteres Base64 a su valor numérico.
var base64Map = func() map[byte]byte {
	m := make(map[byte]byte)
	for i := 0; i < len(base64Table); i++ {
		m[base64Table[i]] = byte(i)
	}
	return m
}()

// Base64ToBytes Función que convierte una cadena en Base64 a un arreglo de bytes.
func Base64ToBytes(encoded string) ([]byte, error) {
	encoded = strings.TrimRight(encoded, "=") // Quitar padding si existe
	if len(encoded)%4 == 1 {
		return nil, errors.New("cadena Base64 inválida")
	}

	var decoded []byte
	var chunk uint32
	count := 0

	// Procesar cada bloque de 4 caracteres
	for i := 0; i < len(encoded); i++ {
		val, exists := base64Map[encoded[i]]
		if !exists {
			return nil, errors.New("carácter inválido en Base64")
		}

		// Acumular 6 bits en un uint32
		chunk = (chunk << 6) | uint32(val)
		count++

		// Extraer bytes de 8 bits cada 4 caracteres procesados
		if count == 4 {
			decoded = append(decoded, byte(chunk>>16), byte(chunk>>8&0xFF), byte(chunk&0xFF))
			chunk = 0
			count = 0
		}
	}

	// Eliminar bytes de relleno
	if mod := len(encoded) % 4; mod > 0 {
		decoded = decoded[:len(decoded)-(3-mod)]
	}

	return decoded, nil
}
