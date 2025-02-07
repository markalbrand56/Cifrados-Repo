package bruteforce

import (
	"Laboratorio-1-B/algorithms"
	"Laboratorio-1-B/scripts/analysis"
	"fmt"
	"strings"
)

// BruteForceVigenere realiza un ataque de fuerza bruta al cifrado de Vigenère.
func BruteForceVigenere(ciphertext string, prefix string, maxLength int) (string, string) {
	bestKey := ""
	bestScore := 0.0
	bestText := ""

	for length := len(prefix); length <= maxLength; length++ {
		nextKeys := GenerateKeys(prefix, length)

		fmt.Println("Probando claves de longitud", length)

		for _, key := range nextKeys {
			decrypted := algorithms.DecryptVigenere(ciphertext, key)
			score := analysis.FrequencyScore(strings.ToLower(decrypted))
			if score > bestScore {
				bestScore = score
				bestKey = key
				bestText = decrypted
			}
		}
	}
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
