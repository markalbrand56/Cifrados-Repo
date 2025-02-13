package main

import (
	"Laboratorio-2-B/algorithms"
	"Laboratorio-2-B/scripts/conversions"
	"Laboratorio-2-B/scripts/images"
	"fmt"
	"image/png"
	"os"
)

func main() {
	fmt.Println("Laboratorio 2-B")

	fmt.Println("1. Cifrado XOR con cadenas de texto")
	text := "Hola"
	key := "abc1"

	fmt.Printf("\t%s XOR %s = %s\n", text, key, algorithms.XOR(conversions.AsciiToBinary(text), conversions.AsciiToBinary(key)))

	fmt.Println("2. Descifrado de imagen con XOR")

	key = "cifrados_2025"

	// Leer archivo y codificar en Base64
	imgData, err := os.ReadFile("examples/imagen_xor.png")
	if err != nil {
		fmt.Println("Error leyendo imagen:", err)
		return
	}

	//encoded := base64.StdEncoding.EncodeToString(imgData)
	encoded := conversions.BytesToBase64(imgData)

	// Decodificar Base64
	//decoded, err := base64.StdEncoding.DecodeString(encoded)
	decoded, err := conversions.Base64ToBytes(encoded)
	if err != nil {
		fmt.Println("Error decodificando base64:", err)
		return
	}

	// Aplicar XOR
	decodedXOR := algorithms.XORBytes(decoded, key)

	// Guardar resultado
	err = os.WriteFile("examples/imagen_descifrada.png", decodedXOR, 0644)
	if err != nil {
		fmt.Println("Error guardando imagen:", err)
	}

	fmt.Println("\tImagen descifrada guardada en examples/imagen_descifrada.png")

	fmt.Println("4. XOR entre dos imágenes")
	file1, err := os.Open("examples/imagen3.png")
	if err != nil {
		fmt.Println("Error leyendo imagen 1:", err)
		return
	}

	file2, err := os.Open("examples/imagen4.png")
	if err != nil {
		fmt.Println("Error leyendo imagen 2:", err)
		return
	}

	img1, err := png.Decode(file1)

	if err != nil {
		fmt.Println("Error leyendo imagen:", err)
		return
	}

	img2, err := png.Decode(file2)

	if err != nil {
		fmt.Println("Error leyendo imagen:", err)
		return
	}

	newImg := images.XORImages(img1, img2)

	out, err := os.Create("examples/imagen_xor_resultado.png")

	defer out.Close()

	if err != nil {
		fmt.Println("Error guardando imagen:", err)
		return
	}

	err = png.Encode(out, newImg)

	if err != nil {
		fmt.Println("Error guardando imagen:", err)
		return
	}

	fmt.Println("\tImagen XOR guardada en examples/imagen_xor_resultado.png")

}
