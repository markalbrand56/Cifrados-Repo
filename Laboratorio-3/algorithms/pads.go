package algorithms

import "errors"

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
