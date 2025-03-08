package main

import (
	"Ejercicio-Block-Cipher/algorithms"
	"Ejercicio-Block-Cipher/scripts"
	"fmt"
	"os"
)

func main() {
	fmt.Println("DES (ECB)")

	file, err := os.ReadFile("inputs/des.txt")

	if err != nil {
		fmt.Println("\nError leyendo archivo:", err)
		return
	}

	fmt.Println("\nTexto plano:", string(file))

	key := scripts.DynamicKey(8)

	fmt.Println("\nClave:", key)

	encryptedDes, err := algorithms.EncryptDESECB(file, key)

	if err != nil {
		fmt.Println("Error cifrando:", err)
		return
	}

	fmt.Println("\nTexto cifrado:", string(encryptedDes))

	decryptedDes, err := algorithms.DecryptDESECB(encryptedDes, key)

	if err != nil {
		fmt.Println("\nError descifrando:", err)
		return
	}

	fmt.Println("\nTexto descifrado:", string(decryptedDes))

}
