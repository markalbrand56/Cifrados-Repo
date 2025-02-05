package bruteforce

import (
	"Laboratorio-1-B/algorithms"
	"Laboratorio-1-B/scripts/analysis"
)

// CaesarBruteForce Realiza un ataque de fuerza bruta para descifrar un texto cifrado con el cifrado César.
func CaesarBruteForce(ciphertext string) (int, string) {
	bestShift := 0
	bestScore := 0.0
	bestText := ""

	for shift := 1; shift <= 30; shift++ {
		decrypted := algorithms.DecryptCaesar(ciphertext, shift)
		score := analysis.FrequencyScore(decrypted)

		if score > bestScore {
			bestScore = score
			bestShift = shift
			bestText = decrypted
		}
	}

	return bestShift, bestText
}
