package analysis

import "strings"

// FrequencyAnalysis Analiza la frecuencia de los caracteres en un texto. Retorna un mapa con la frecuencia de cada letra.
func FrequencyAnalysis(text string) map[rune]float64 {
	text = strings.ToUpper(text)
	frequency := make(map[rune]float64)
	total := 0

	for _, c := range text {
		if c >= 'A' && c <= 'Z' {
			frequency[c]++
			total++
		}
	}

	for c := 'A'; c <= 'Z'; c++ {
		frequency[c] /= float64(total)
	}

	return frequency
}
