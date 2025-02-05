package analysis

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"log"
	"strings"
)

var spanishFrequency = map[rune]float64{
	'A': 0.1253,
	'B': 0.0142,
	'C': 0.0468,
	'D': 0.0586,
	'E': 0.1368,
	'F': 0.0069,
	'G': 0.0101,
	'H': 0.0070,
	'I': 0.0625,
	'J': 0.0044,
	'K': 0.0002,
	'L': 0.0497,
	'M': 0.0315,
	'N': 0.0671,
	'Ñ': 0.0031,
	'O': 0.0868,
	'P': 0.0251,
	'Q': 0.0088,
	'R': 0.0687,
	'S': 0.0798,
	'T': 0.0463,
	'U': 0.0393,
	'V': 0.0090,
	'W': 0.0002,
	'X': 0.0022,
	'Y': 0.0090,
	'Z': 0.0052,
}

// FrequencyAnalysis Analiza la frecuencia de los caracteres en un texto. Retorna un mapa con la frecuencia de cada letra.
func FrequencyAnalysis(text string) map[rune]float64 {
	text = strings.ToUpper(text) // Convertir a mayúsculas
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

func CompareFrequency(frequency map[rune]float64) {
	// Crear un gráfico
	p := plot.New()
	p.Title.Text = "Comparación de Frecuencias de Letras"
	p.X.Label.Text = "Letras"
	p.Y.Label.Text = "Frecuencia"

	// Listas de valores para graficar
	letras := "ABCDEFGHIJKLMNOPQRSTUVWXYZÑ"
	n := len(letras)

	espData := make(plotter.Values, n)
	textData := make(plotter.Values, n)

	for i, c := range letras {
		espData[i] = spanishFrequency[c]
		textData[i] = frequency[c]
	}

	// Crear barras
	barWidth := vg.Points(8)

	barsEsp, err := plotter.NewBarChart(espData, barWidth)
	if err != nil {
		log.Fatal(err)
	}
	barsEsp.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Rojo

	barsText, err := plotter.NewBarChart(textData, barWidth)
	if err != nil {
		log.Fatal(err)
	}
	barsText.Color = color.RGBA{B: 255, A: 255} // Azul

	// Ajustar posiciones de las barras para que no se sobrepongan
	barsEsp.Offset = -barWidth / 2
	barsText.Offset = barWidth / 2

	// Agregar barras al gráfico
	p.Add(barsEsp, barsText)

	p.Legend.Add("Español", barsEsp)
	p.Legend.Add("Texto", barsText)
	p.Legend.Top = true

	// Etiquetas en el eje X
	p.NominalX(strings.Split(letras, "")...)

	// Guardar la imagen
	if err := p.Save(12*vg.Inch, 6*vg.Inch, "out/frecuencia.png"); err != nil {
		log.Fatal(err)
	}
}
