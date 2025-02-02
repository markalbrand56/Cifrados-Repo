package scripts

import (
	"crypto/rand"
	"fmt"
)

// DynamicKey Genera una clave aleatoria de la longitud especificada.
func DynamicKey(length int) string {
	if length <= 0 {
		fmt.Println("La longitud de la clave debe ser mayor a 0")
		return ""
	}

	const asciiStart = 33 // Carácter ASCII más bajo ('!')
	const asciiEnd = 126  // Carácter ASCII más alto ('~')

	key := make([]byte, length)
	for i := 0; i < length; i++ {
		// Genera un valor aleatorio dentro del rango ASCII imprimible
		randomByte := make([]byte, 1)
		_, err := rand.Read(randomByte) // bytes aleatorios criptográficamente seguros

		if err != nil {
			fmt.Println("Error al generar el valor aleatorio")
			return ""
		}

		// Ajusta el valor para que esté en el rango imprimible
		key[i] = byte(asciiStart + (int(randomByte[0]) % (asciiEnd - asciiStart + 1)))
	}

	return string(key)
}
