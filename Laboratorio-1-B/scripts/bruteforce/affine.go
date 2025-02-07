package bruteforce

import (
	"Laboratorio-1-B/algorithms"
	"Laboratorio-1-B/scripts/analysis"
	"fmt"
	"strings"
)

// AffineBruteForce realiza un ataque de fuerza bruta al cifrado afín.
func AffineBruteForce(ciphertext string) (int, int, string) {
	bestA, bestB := 1, 1
	bestScore := 0.0
	bestText := ""
	validAs := []int{1, 2, 4, 5, 7, 8, 10, 11, 13, 14, 16, 17, 19, 20, 22, 23, 25, 26}

	for _, a := range validAs {
		for b := 1; b <= 16; b++ {
			decrypted := algorithms.DecryptAffine(ciphertext, a, b)

			fmt.Printf("a = %d, b = %d: %s\n", a, b, decrypted)

			if decrypted == "" {
				continue
			}
			score := analysis.FrequencyScore(strings.ToLower(decrypted))

			if score > bestScore {
				bestScore = score
				bestA = a
				bestB = b
				bestText = decrypted
			}
		}
	}

	fmt.Printf("El mejor puntaje fue: %f\n", bestScore)

	return bestA, bestB, bestText
}
