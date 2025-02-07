package main

import (
	"Laboratorio-1-B/scripts/bruteforce"
	"fmt"
	"os"
)

// filter Elimina los caracteres no alfabéticos de un texto.
func filter(text string) string {
	var filteredText []rune

	for _, c := range text {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			filteredText = append(filteredText, c)
		}
	}

	return string(filteredText)
}

func main() {
	file, err := os.ReadFile("examples/ceasar.txt")

	if err != nil {
		panic(err)
	}

	encryptedCaesar := string(file)

	// Imprimir el contenido del archivo
	fmt.Printf("\nCifrado César:\n%s\n", encryptedCaesar)

	// Fuerza bruta para descifrar el texto
	shift, decryptedCaesar := bruteforce.CaesarBruteForce(encryptedCaesar)

	// Imprimir el texto descifrado
	fmt.Printf("\nDescifrado César (desplazamiento %d):\n%s\n", shift, decryptedCaesar)

	file, err = os.ReadFile("examples/afines.txt")

	encryptedAffine := string(file)

	fmt.Printf("\nCifrado Afín:\n%s\n", filter(encryptedAffine))

	a, b, decryptedAffine := bruteforce.AffineBruteForce(filter(encryptedAffine))

	fmt.Printf("\nDescifrado Afín (a = %d, b = %d):\n%s\n", a, b, decryptedAffine)

	file, err = os.ReadFile("examples/vigenere.txt")

	encryptedVigenere := string(file)

	fmt.Printf("\nCifrado Vigenère:\n%s\n", filter(encryptedVigenere))

	key, decryptedVigenere := bruteforce.VigenereBruteForce(encryptedVigenere, "PA", 6)

	fmt.Printf("\nDescifrado Vigenère (clave = %s):\n%s\n", key, decryptedVigenere)
}
