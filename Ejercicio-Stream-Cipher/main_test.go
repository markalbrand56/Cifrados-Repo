package main

import (
	"Ejercicio-Stream-Cipher/algorithms"
	"Ejercicio-Stream-Cipher/scripts/keys"
	"slices"
	"testing"
)

func TestCifrar(t *testing.T) {
	message := "Hello, world!"
	seed := uint32(1234567890)

	keystream := keys.GenerateKeystream(len(message), seed)

	ciphertext := algorithms.XORBytes([]byte(message), string(keystream))

	plaintext := algorithms.XORBytes(ciphertext, string(keystream))

	if message != string(plaintext) {
		t.Errorf("Error en la función XORBytes")
	}
}

func TestGenerarKeystream(t *testing.T) {
	length := 10
	seed := uint32(1234567890)

	keystream := keys.GenerateKeystream(length, seed)

	if len(keystream) != length {
		t.Errorf("Error en la función GenerateKeystream")
	}

	if slices.Equal(keystream, keys.GenerateKeystream(length, seed-1)) {
		t.Errorf("Error en la función GenerateKeystream, debería ser diferente por la aleatoriedad")
	}
}

func TestXORBytes(t *testing.T) {
	data := []byte{0x00, 0x01, 0x02, 0x03}
	key := "key"

	result := algorithms.XORBytes(data, key)

	if len(result) != len(data) {
		t.Errorf("Error en la función XORBytes")
	}
}

func TestXOR(t *testing.T) {
	a := "1010101010101010"
	b := "0101010101010101"

	result := algorithms.XOR(a, b)

	if len(result) != len(a) {
		t.Errorf("Error en la función XOR")
	}
}
