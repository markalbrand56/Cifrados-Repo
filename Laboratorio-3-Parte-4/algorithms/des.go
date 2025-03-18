package algorithms

import (
	"crypto/des"
)

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
