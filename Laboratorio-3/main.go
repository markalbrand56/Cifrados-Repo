package main

import (
	"Laboratorio-3/algorithms"
	"Laboratorio-3/scripts"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Parte 1: Rompiendo ECB en imágenes")

	// Leer la imagen
	//filename := "tux.ppm"
	filename := "pexels-rodolfoclix-922610.ppm"

	image, err := os.ReadFile(fmt.Sprintf("inputs/%s", filename))

	if err != nil {
		fmt.Println("\nError leyendo archivo:", err)
		return
	}

	// Separar el encabezado de los datos de píxeles
	headerEnd := findHeaderEnd(image)
	if headerEnd == -1 {
		fmt.Println("Formato de imagen PPM no reconocido")
		return
	}

	header := image[:headerEnd]
	pixelData := image[headerEnd:]

	// AES ECB
	key := scripts.DynamicKey(16)
	fmt.Println("\nClave:", string(key))

	// Cifrar solo los datos de píxeles con AES-ECB
	encryptedPixels, err := algorithms.EncryptAESECB(pixelData, key)
	if err != nil {
		fmt.Println("Error cifrando:", err)
		return
	}

	// Guardar la imagen cifrada (manteniendo el encabezado intacto)
	err = os.WriteFile(fmt.Sprintf("./outputs/aes_ecb_%s", filename), append(header, encryptedPixels...), 0644)

	if err != nil {
		fmt.Println("Error escribiendo archivo cifrado:", err)
		return
	}

	fmt.Println("Imagen cifrada correctamente con AES-ECB")

	// AES CBC

	// Generar IV aleatorio
	ivAES := scripts.DynamicKey(16)
	fmt.Println("\nIV:", string(ivAES))

	// Cifrar solo los datos de píxeles con AES-CBC
	encryptedPixels, err = algorithms.EncryptAESCBC(pixelData, key, ivAES)

	if err != nil {
		fmt.Println("Error cifrando:", err)
		return
	}

	// Guardar la imagen cifrada (manteniendo el encabezado intacto)
	err = os.WriteFile(fmt.Sprintf("./outputs/aes_cbc_%s", filename), append(header, encryptedPixels...), 0644)
}

// Encuentra el fin del encabezado PPM
func findHeaderEnd(data []byte) int {
	headerStr := string(data)
	index := strings.Index(headerStr, "\n255\n")

	if index != -1 {
		return index + len("\n255\n")
	}

	return -1
}
