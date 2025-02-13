package main

import (
	"Laboratorio-2-A/algorithms"
	"Laboratorio-2-A/scripts/conversions"
	"fmt"
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
	fmt.Println("Laboratorio-2-A")

	fmt.Println("1. Cadenas a bytes")

	fmt.Printf("%s: %v\n", "Hola", conversions.AsciiToBinary("Hola"))
	fmt.Printf("%s: %v\n", "Adios", conversions.AsciiToBinary("Adios"))

	fmt.Println("2. Bytes a cadenas")

	fmt.Printf("%s: %s\n", "01001000 01101111 01101100 01100001", conversions.BinaryToAscii("01001000 01101111 01101100 01100001"))
	fmt.Printf("%s: %s\n", "01000001 01100100 01101001 01101111 01110011", conversions.BinaryToAscii("01000001 01100100 01101001 01101111 01110011"))

	fmt.Println("3. Texto a base64")

	fmt.Printf("%s: %s\n", "Hola", conversions.AsciiToBase64("Hola"))
	fmt.Printf("%s: %s\n", "Adios", conversions.AsciiToBase64("Adios"))

	fmt.Println("4. Base64 a texto")

	fmt.Printf("%s: %s\n", "SG9sYQ==", conversions.Base64ToAscii("SG9sYQ=="))
	fmt.Printf("%s: %s\n", "QWRpb3M=", conversions.Base64ToAscii("QWRpb3M="))

	fmt.Println("5. XOR entre dos cadenas")

	binary1 := conversions.AsciiToBinary(filter("Hola"))
	binary2 := conversions.AsciiToBinary(filter("Adios"))

	fmt.Printf("%s (%s) XOR %s (%s): %s\n", binary1, "Hola", binary2, "Adios", algorithms.XOR(binary1, binary2))

}
