package main

import (
	"Ejercicio-2/algorithms"
	"fmt"
)

func main() {
	var word = "Hello World"
	//var base64 = "SGVsbG8gd29ybGQ=" // Hello world en base64

	fmt.Printf("Texto: %s\n", word)

	caesar := algorithms.Caesar(word, 3)
	fmt.Printf("\nCifrado César: %s\n", caesar)

	affine := algorithms.Affine(word, 3, 5)
	fmt.Printf("\nCifrado Afín: %s\n", affine)
	fmt.Printf("Descifrado Afín: %s\n", algorithms.DecryptAffine(affine, 3, 5))
}
