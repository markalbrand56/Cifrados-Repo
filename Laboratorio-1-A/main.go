package main

import (
	"Laboratorio-1-A/algorithms"
	"fmt"
)

func main() {
	var word = "Hello World"
	//var base64 = "SGVsbG8gd29ybGQ=" // Hello world en base64

	fmt.Printf("Texto: %s\n", word)

	caesar := algorithms.Caesar(word, 3)
	fmt.Printf("\nCifrado César: %s\n", caesar)
	fmt.Printf("Descifrado César: %s\n", algorithms.DecryptCaesar(caesar, 3))

	affine := algorithms.Affine(word, 3, 5)
	fmt.Printf("\nCifrado Afín: %s\n", affine)
	fmt.Printf("Descifrado Afín: %s\n", algorithms.DecryptAffine(affine, 3, 5))

	vigenere := algorithms.Vigenere(word, "key")
	fmt.Printf("\nCifrado Vigenère: %s\n", vigenere)
	fmt.Printf("Descifrado Vigenère: %s\n", algorithms.DecryptVigenere(vigenere, "key"))
}
