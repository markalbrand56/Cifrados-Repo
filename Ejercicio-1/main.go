package main

import (
	"Ejercicio-1/scripts"
	"fmt"
	"strings"
)

func main() {
	var word = "Hello World"

	fmt.Printf("Texto: %s\n", word)

	binary := scripts.AsciiToBinary(word)

	fmt.Printf("ASCII a binario: %s\n", binary)

	base64 := "SGVsbG8gd29ybGQ="

	fmt.Printf("Base64 a binario: %s\n", scripts.Base64ToBinary(base64))

	fmt.Printf("Binario a base 64: %s\n", scripts.BinaryToBase64(strings.Replace(binary, " ", "", -1)))

	fmt.Printf("Binario a ASCII: %s\n", scripts.BinaryToAscii(binary))

	fmt.Printf("Base64 a ASCII: %s\n", scripts.Base64ToAscii(base64))

	fmt.Printf("ASCII a Base64: %s\n", scripts.AsciiToBase64(word))

}
