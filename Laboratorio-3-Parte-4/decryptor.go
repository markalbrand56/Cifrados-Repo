package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"Laboratorio-3-Parte-4/algorithms"
)

func findEncryptedFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".locked") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func decryptFiles(files []string, key []byte) {
	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error leyendo archivo:", file, err)
			continue
		}

		encryptedData := data[16:]

		decryptedData, err := algorithms.DecryptAESCBC(encryptedData, key)
		if err != nil {
			fmt.Println("Error descifrando archivo:", file, err)
			continue
		}

		originalFile := strings.TrimSuffix(file, ".locked")
		err = ioutil.WriteFile(originalFile, decryptedData, 0644)
		if err != nil {
			fmt.Println("Error escribiendo archivo descifrado:", originalFile, err)
		} else {
			fmt.Println("Archivo descifrado:", originalFile)
			os.Remove(file) // Elimina el archivo cifrado
		}
	}
}

func main() {
	fmt.Print("Ingresa la clave de descifrado: ")
	var keyHex string
	fmt.Scanln(&keyHex)

	key, err := hex.DecodeString(keyHex)
	if err != nil || len(key) != 32 {
		fmt.Println("Clave inválida.")
		return
	}

	rootDir := "."
	fmt.Println("Buscando archivos cifrados en:", rootDir)

	files, err := findEncryptedFiles(rootDir)
	if err != nil {
		fmt.Println("Error buscando archivos cifrados:", err)
		return
	}

	if len(files) == 0 {
		fmt.Println("No se encontraron archivos cifrados.")
		return
	}

	fmt.Println("\nArchivos cifrados encontrados:")
	for _, file := range files {
		fmt.Println("-", file)
	}

	fmt.Print("\n¿Deseas continuar con el descifrado? (escribe 'YES' para confirmar): ")
	var confirm string
	fmt.Scanln(&confirm)
	if strings.ToUpper(confirm) != "YES" {
		fmt.Println("Operación cancelada.")
		return
	}

	decryptFiles(files, key)
	fmt.Println("Proceso de descifrado finalizado.")
}
