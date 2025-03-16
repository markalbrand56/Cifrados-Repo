package algorithms

import (
	"crypto/aes"
	"crypto/cipher"
)

// EncryptAESECB Cifrado AES en modo ECB
func EncryptAESECB(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	paddedText := padPKCS5(plainText, block.BlockSize())
	cipherText := make([]byte, len(paddedText))
	for i := 0; i < len(paddedText); i += block.BlockSize() {
		block.Encrypt(cipherText[i:i+block.BlockSize()], paddedText[i:i+block.BlockSize()])
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

// EncryptAESCBC Cifrado AES en modo CBC
func EncryptAESCBC(plainText, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	paddedText := padPKCS5(plainText, block.BlockSize())
	cipherText := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, paddedText)
	return cipherText, nil
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
