package main

import (
	"Ejercicio-Block-Cipher/algorithms"
	"testing"
)

func TestDes(t *testing.T) {
	input := "Hello, world!"

	key := "12345678"

	encryptedDes, err := algorithms.EncryptDESECB([]byte(input), []byte(key))

	if err != nil {
		t.Error("Error cifrando: ", err)
	}

	decryptedDes, err := algorithms.DecryptDESECB(encryptedDes, []byte(key))

	if err != nil {
		t.Error("Error descifrando: ", err)
	}

	if input != string(decryptedDes) {
		t.Error("Error en el cifrado DES")
	}
}

func TestTripleDes(t *testing.T) {
	input := "Hello, world!"

	key := "987654321234567898765432"

	iv := "12345678"

	encrypted3Des, err := algorithms.Encrypt3DESCBC([]byte(input), []byte(key), []byte(iv))

	if err != nil {
		t.Error("Error cifrando: ", err)
	}

	decrypted3Des, err := algorithms.Decrypt3DESCBC(encrypted3Des, []byte(key), []byte(iv))

	if err != nil {
		t.Error("Error descifrando: ", err)
	}

	if input != string(decrypted3Des) {
		t.Error("Error en el cifrado 3DES")
	}
}

func TestAes(t *testing.T) {
	input := "Hello, world!"

	key := "1234567890123456" // Clave de 16 bytes para AES-128
	iv := "1234567890123456"  // IV de 16 bytes para AES-CBC

	// Prueba con AES-CBC
	encryptedAes, err := algorithms.EncryptAESCBC([]byte(input), []byte(key), []byte(iv))
	if err != nil {
		t.Error("Error cifrando: ", err)
	}

	decryptedAes, err := algorithms.DecryptAESCBC(encryptedAes, []byte(key), []byte(iv))
	if err != nil {
		t.Error("Error descifrando: ", err)
	}

	if input != string(decryptedAes) {
		t.Error("Error en el cifrado AES")
	}

	// Prueba con AES-ECB
	encryptedAes, err = algorithms.EncryptAESECB([]byte(input), []byte(key))
	if err != nil {
		t.Error("Error cifrando AES-ECB: ", err)
	}

	decryptedAes, err = algorithms.DecryptAESECB(encryptedAes, []byte(key))
	if err != nil {
		t.Error("Error descifrando AES-ECB: ", err)
	}

	if input != string(decryptedAes) {
		t.Error("Error en el cifrado AES-ECB")
	}
}
