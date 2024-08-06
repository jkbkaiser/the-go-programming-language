package main

import (
    "fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
    "strconv"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

var (
    cycles  = 5    // number of complete x oscillator revolutions
    res     = 0.001 // angular resolution
    size    = 100   // image canvase covers [-size..+size]
    nframes = 64    // number of animation frames
    delay   = 8     // delay between frames in 10ms units
)


func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        log.Print(err)
    }
    
    passed_cycles_str := r.Form.Get("cycles")

    if len(passed_cycles_str) != 0 {
        passed_cycles_int, err := strconv.Atoi(passed_cycles_str)

        if err != nil {
            cycles = passed_cycles_int
        }
    }

    lissajous(w)
}

func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size-1, 2*size-1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
