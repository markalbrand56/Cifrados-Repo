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

	// 3DES
	fmt.Println("\n3DES (CBC)")

	file, err = os.ReadFile("inputs/3des.txt")

	if err != nil {
		fmt.Println("\nError leyendo archivo:", err)
		return
	}

	fmt.Println("\nTexto plano:", string(file))

	key = scripts.DynamicKey(24)

	fmt.Println("\nClave:", key)

	iv := scripts.DynamicKey(8)

	fmt.Println("\nIV:", iv)

	encrypted3Des, err := algorithms.Encrypt3DESCBC(file, key, iv)

	if err != nil {
		fmt.Println("Error cifrando:", err)
		return
	}

	fmt.Println("\nTexto cifrado:", string(encrypted3Des))

	decrypted3Des, err := algorithms.Decrypt3DESCBC(encrypted3Des, key, iv)

	if err != nil {
		fmt.Println("\nError descifrando:", err)
		return
	}

	fmt.Println("\nTexto descifrado:", string(decrypted3Des))

	// AES (CBC y ECB)

	fmt.Println("\nAES (CBC)")

	image, err := os.ReadFile("inputs/pic.png")

	if err != nil {
		fmt.Println("\nError leyendo archivo:", err)
		return
	}

	key = scripts.DynamicKey(16)

	fmt.Println("\nClave:", key)

	ivAES := scripts.DynamicKey(16) // Inicialización de vector aleatorio

	fmt.Println("\nIV:", ivAES)

	encryptedAes, err := algorithms.EncryptAESCBC(image, key, ivAES)

	if err != nil {
		fmt.Println("Error cifrando:", err)
		return
	}

	fmt.Println("\nImagen cifrada:", encryptedAes)
	err = os.WriteFile("./outputs/pic_encrypted_aes.png", encryptedAes, 0644)

	decryptedAes, err := algorithms.DecryptAESCBC(encryptedAes, key, ivAES)

	if err != nil {
		fmt.Println("\nError descifrando:", err)
		return
	}

	err = os.WriteFile("./outputs/pic_descifrada_aes.png", decryptedAes, 0644)
	if err != nil {
		fmt.Println("Error guardando imagen:", err)
		return
	}

	fmt.Println("\nAES (ECB)")

	encryptedAes, err = algorithms.EncryptAESECB(image, key)

	if err != nil {
		fmt.Println("Error cifrando:", err)
		return
	}

	fmt.Println("\nImagen cifrada:", encryptedAes)
	err = os.WriteFile("./outputs/pic_encrypted_aes_ecb.png", encryptedAes, 0644)

	decryptedAes, err = algorithms.DecryptAESECB(encryptedAes, key)

	if err != nil {
		fmt.Println("\nError descifrando:", err)
		return
	}

	err = os.WriteFile("./outputs/pic_descifrada_aes_ecb.png", decryptedAes, 0644)

	if err != nil {
		fmt.Println("Error guardando imagen:", err)
		return
	}
}
