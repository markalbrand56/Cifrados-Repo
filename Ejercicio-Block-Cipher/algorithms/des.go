package algorithms

import (
	"crypto/des"
	"errors"
)

// Padding manual para DES (PKCS5)
func padPKCS5(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	paddedData := append(data, make([]byte, padding)...)
	for i := len(data); i < len(paddedData); i++ {
		paddedData[i] = byte(padding)
	}
	return paddedData
}

// Remueve el padding PKCS5
func unpadPKCS5(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("data is empty")
	}
	padding := int(data[len(data)-1])
	if padding > len(data) {
		return nil, errors.New("invalid padding")
	}
	return data[:len(data)-padding], nil
}

// EncryptDESECB Cifrado DES en modo ECB
func EncryptDESECB(plainText, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
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

// DecryptDESECB Descifrado DES en modo ECB
func DecryptDESECB(cipherText, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainText := make([]byte, len(cipherText))
	for i := 0; i < len(cipherText); i += block.BlockSize() {
		block.Decrypt(plainText[i:i+block.BlockSize()], cipherText[i:i+block.BlockSize()])
	}
	return unpadPKCS5(plainText)
}
