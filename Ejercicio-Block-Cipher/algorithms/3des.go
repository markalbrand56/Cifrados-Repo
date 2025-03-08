package algorithms

import (
	"crypto/cipher"
	"crypto/des"
)

// Encrypt3DESCBC Cifrado 3DES en modo CBC
func Encrypt3DESCBC(plainText, key, iv []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	paddedText := padPKCS5(plainText, block.BlockSize())
	cipherText := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, paddedText)
	return cipherText, nil
}

// Decrypt3DESCBC Descifrado 3DES en modo CBC
func Decrypt3DESCBC(cipherText, key, iv []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plainText, cipherText)
	return unpadPKCS5(plainText)
}
