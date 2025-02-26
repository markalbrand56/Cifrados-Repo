package main

import (
	"Ejercicio-Stream-Cipher/algorithms"
	"Ejercicio-Stream-Cipher/scripts/keys"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var message string
	var seed uint32
	var keystream []byte

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter message:")
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	message = message[:len(message)-1]

	fmt.Println("Enter seed:")
	_, err = fmt.Scanln(&seed)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	keystream = keys.GenerateKeystream(len(message), seed)

	fmt.Println("Keystream:", keystream)

	ciphertext := algorithms.XORBytes([]byte(message), string(keystream))

	fmt.Println("Ciphertext:", ciphertext)

	plaintext := algorithms.XORBytes(ciphertext, string(keystream))

	fmt.Println("Plaintext:", string(plaintext))
}
