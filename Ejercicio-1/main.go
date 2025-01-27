package main

import (
	"Ejercicio-1/algorithms"
	"Ejercicio-1/scripts"
	"fmt"
)

func main() {
	var word = "Hello World"

	fmt.Printf("Texto: %s\n", word)

	binary := scripts.AsciiToBinary(word)

	fmt.Printf("ASCII a binario: %s\n", binary)

	base64 := "SGVsbG8gd29ybGQ="

	fmt.Printf("Base64 a binario: %s\n", scripts.Base64ToBinary(base64))

	fmt.Printf("Binario a base 64: %s\n", scripts.BinaryToBase64(binary))

	fmt.Printf("Binario a ASCII: %s\n", scripts.BinaryToAscii(binary))

	fmt.Printf("Base64 a ASCII: %s\n", scripts.Base64ToAscii(base64))

	fmt.Printf("ASCII a Base64: %s\n", scripts.AsciiToBase64(word))

	a := "10101010"
	b := "01010101"

	fmt.Printf("XOR de %s y %s: %s\n", a, b, algorithms.XOR(a, b)) // 11111111

	fmt.Printf("Llave dinámica: %s\n", scripts.DynamicKey(15))

	fixed := algorithms.FixedCypher("abc", "123")

	fmt.Printf("Cypher fijo: %s\n", fixed)
	fmt.Printf("Descifrado (fijo): %s\n", algorithms.DecryptCypher(fixed, "123"))

	dynamicKey := scripts.DynamicKey(15)
	dynamic := algorithms.DynamicCypher(word, dynamicKey)

	fmt.Printf("Cypher dinámico: %s\n", dynamic)
	fmt.Printf("Descifrado (dinámico): %s\n", algorithms.DecryptCypher(dynamic, dynamicKey))

}
