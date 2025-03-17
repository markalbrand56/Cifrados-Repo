package algorithms

import (
	"crypto/aes"
	"crypto/cipher"
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

// DecryptAESECB Descifrado AES en modo ECB
func DecryptAESECB(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText := make([]byte, len(cipherText))
	for i := 0; i < len(cipherText); i += block.BlockSize() {
		block.Decrypt(plainText[i:i+block.BlockSize()], cipherText[i:i+block.BlockSize()])
	}
	return unpadPKCS5(plainText)
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

// DecryptAESCBC Descifrado AES en modo CBC
func DecryptAESCBC(cipherText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)
	return unpadPKCS5(plainText)
}
