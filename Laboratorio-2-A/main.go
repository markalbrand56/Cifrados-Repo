package main

import "fmt"

// filter Elimina los caracteres no alfabéticos de un texto.
func filter(text string) string {
	var filteredText []rune

	for _, c := range text {
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			filteredText = append(filteredText, c)
		}
	}

	return string(filteredText)
}

func main() {
	fmt.Println("Laboratorio-2-A")
}
