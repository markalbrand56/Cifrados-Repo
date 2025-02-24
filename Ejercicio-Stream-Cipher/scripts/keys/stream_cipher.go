package keys

import "Ejercicio-Stream-Cipher/scripts/random"

// GenerateKeystream genera un keystream pseudoaleatorio basado en una semilla
func GenerateKeystream(length int, seed uint32) []byte {
	lcg := random.NewLCG(seed)
	keystream := make([]byte, length)

	for i := 0; i < length; i++ {
		keystream[i] = lcg.Next()
	}

	return keystream
}
