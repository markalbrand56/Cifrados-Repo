package bruteforce

import (
	"Laboratorio-2-A/algorithms"
	"Laboratorio-2-A/scripts/analysis"
	"strings"
)

// VigenereBruteForce realiza un ataque de fuerza bruta al cifrado de Vigenère.
func VigenereBruteForce(ciphertext string, prefix string, maxLength int) (string, string) {
	bestKey := ""
	bestScore := 0.0
	bestText := ""
	letters := "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ"

	var generate func(string, int)
	generate = func(currentKey string, length int) {
		if length == maxLength {
			decrypted := algorithms.DecryptVigenere(ciphertext, currentKey)
			score := analysis.FrequencyScore(strings.ToLower(decrypted))
			if score > bestScore {
				bestScore = score
				bestKey = currentKey
				bestText = decrypted
			}
			return
		}
		for _, char := range letters {
			generate(currentKey+string(char), length+1)
		}
	}

	generate(prefix, len(prefix))
	return bestKey, bestText
}

// GenerateKeys genera todas las posibles claves con el prefijo dado y longitud máxima.
func GenerateKeys(prefix string, length int) []string {
	letters := "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ"
	keys := []string{}
	if len(prefix) == length {
		return []string{prefix}
	}
	for _, char := range letters {
		keys = append(keys, GenerateKeys(prefix+string(char), length)...)
	}
	return keys
}
