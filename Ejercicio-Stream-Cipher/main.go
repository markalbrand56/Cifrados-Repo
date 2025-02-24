package main

import (
	"Ejercicio-Stream-Cipher/algorithms"
	"Ejercicio-Stream-Cipher/scripts/keys"
	"fmt"
)

func main() {
	message := "Hello, world!"
	seed := uint32(1234567890)
	keystream := keys.GenerateKeystream(len(message), seed)

	ciphertext := algorithms.XORBytes([]byte(message), string(keystream))
	fmt.Println("Ciphertext:", ciphertext)

	plaintext := algorithms.XORBytes(ciphertext, string(keystream))
	fmt.Println("Plaintext:", string(plaintext))
}
