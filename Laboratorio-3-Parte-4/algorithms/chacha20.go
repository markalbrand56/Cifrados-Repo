package algorithms

import (
	"crypto/rand"
	"golang.org/x/crypto/chacha20"
	"io"
)

// GenerateNonce genera un nonce de 12 bytes para ChaCha20
func GenerateNonce() ([]byte, error) {
	nonce := make([]byte, chacha20.NonceSize)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

// EncryptChaCha20 cifra un mensaje con ChaCha20
func EncryptChaCha20(plainText, key []byte) ([]byte, []byte, error) {
	nonce, err := GenerateNonce()
	if err != nil {
		return nil, nil, err
	}

	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, nil, err
	}

	cipherText := make([]byte, len(plainText))
	cipher.XORKeyStream(cipherText, plainText)

	return cipherText, nonce, nil
}

// DecryptChaCha20 descifra un mensaje cifrado con ChaCha20
func DecryptChaCha20(cipherText, key, nonce []byte) ([]byte, error) {
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	plainText := make([]byte, len(cipherText))
	cipher.XORKeyStream(plainText, cipherText)

	return plainText, nil
}
