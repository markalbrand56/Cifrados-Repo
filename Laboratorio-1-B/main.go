package main

import (
	"Laboratorio-1-B/scripts/bruteforce"
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("examples/ceasar.txt")

	if err != nil {
		panic(err)
	}

	encryptedCaesar := string(file)

	// Imprimir el contenido del archivo
	fmt.Printf("Cifrado César:\n%s\n", encryptedCaesar)

	// Fuerza bruta para descifrar el texto
	shift, decryptedCaesar := bruteforce.CaesarBruteForce(encryptedCaesar)

	// Imprimir el texto descifrado
	fmt.Printf("Descifrado César (desplazamiento %d):\n%s\n", shift, decryptedCaesar)
}
