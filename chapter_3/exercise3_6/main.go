package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			a1, a2, a3 := mandelbrot(complex(x+0.5, y+0.5))
			b1, b2, b3 := mandelbrot(complex(x-0.5, y+0.5))
			c1, c2, c3 := mandelbrot(complex(x+0.5, y-0.5))
			d1, d2, d3 := mandelbrot(complex(x-0.5, y-0.5))

			c := color.RGBA{
                (a1+b1+c1+d1)/4,
                (a2+b2+c2+d2)/4,
                (a3+b3+c3+d3)/4,
                100,
            }

			// Image point (px, py) represents complex value z.
			img.Set(px, py, c)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) (uint8, uint8, uint8) {
	const iterations = 200
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return getColor(n)
		}
	}

	return 255, 255, 255
}

func getColor(n uint8) (uint8, uint8, uint8) {
	return (n * 17) % 128, (n * 29) % 128, (n * 73) % 128
}
