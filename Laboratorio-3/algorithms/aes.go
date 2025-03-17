package algorithms

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// EncryptAESECB cifra los datos en modo ECB con padding
func EncryptAESECB(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Aplicar padding si es necesario
	paddedText := padPKCS7(plainText, aes.BlockSize)

	cipherText := make([]byte, len(paddedText))
	for i := 0; i < len(paddedText); i += aes.BlockSize {
		block.Encrypt(cipherText[i:i+aes.BlockSize], paddedText[i:i+aes.BlockSize])
	}

	return cipherText, nil
}

// DecryptAESECB descifra AES en modo ECB
func DecryptAESECB(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(cipherText))
	for i := 0; i < len(cipherText); i += block.BlockSize() {
		block.Decrypt(plainText[i:i+block.BlockSize()], cipherText[i:i+block.BlockSize()])
	}

	return unpadPKCS7(plainText)
}

// EncryptAESCBC cifra los datos en modo CBC con padding
func EncryptAESCBC(plainText, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Aplicar padding si es necesario
	paddedText := padPKCS7(plainText, aes.BlockSize)

	cipherText := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, paddedText)

	// Devolver IV + datos cifrados (el IV debe almacenarse para el descifrado)
	return append(iv, cipherText...), nil
}

// DecryptAESCBC descifra AES en modo CBC
func DecryptAESCBC(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// **Corrección**: Extraer el IV antes de descifrar
	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("ciphertext demasiado corto")
	}
	iv := cipherText[:aes.BlockSize]            // Extraer IV
	encryptedData := cipherText[aes.BlockSize:] // Extraer datos cifrados

	// **IMPORTANTE**: Asegurar que el tamaño sea un múltiplo del bloque
	if len(encryptedData)%aes.BlockSize != 0 {
		return nil, errors.New("ciphertext no es múltiplo del tamaño del bloque")
	}

	plainText := make([]byte, len(encryptedData))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, encryptedData)

	return unpadPKCS7(plainText)
}
