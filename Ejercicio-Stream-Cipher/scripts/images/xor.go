package images

import (
	"image"
	"image/color"
)

// XORImages aplica XOR pixel a pixel entre dos imágenes
func XORImages(img1, img2 image.Image) *image.RGBA {
	bounds := img1.Bounds()
	newImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c1 := img1.At(x, y).(color.RGBA)
			c2 := img2.At(x, y).(color.RGBA)

			newImg.Set(x, y, color.RGBA{
				R: c1.R ^ c2.R,
				G: c1.G ^ c2.G,
				B: c1.B ^ c2.B,
				A: 255, // Mantener transparencia
			})
		}
	}

	return newImg
}
