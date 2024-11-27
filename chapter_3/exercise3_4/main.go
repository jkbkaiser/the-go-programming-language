package main

// Example query http://localhost:8000/?width=1000&height=600&peak_color=%230000FF&valley_color=%230000FF

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = float64(width) / 2 / float64(xyrange) // pixels per x or y unit
	zscale        = float64(height) * 0.4
	angle         = math.Pi / 6
	valley_color  = "#0000ff"
	peak_color    = "#ff0000"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func writeSvg(w http.ResponseWriter) {
	fmt.Printf(peak_color)
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			if math.IsInf(ax, 0) ||
				math.IsInf(ay, 0) ||
				math.IsInf(bx, 0) ||
				math.IsInf(by, 0) ||
				math.IsInf(cx, 0) ||
				math.IsInf(cy, 0) ||
				math.IsInf(dx, 0) ||
				math.IsInf(dy, 0) {
				continue
			}
			color := valley_color
			if (az+bz+cz+dz)/4 > 0 {
				color = peak_color
			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*float64(xyscale)
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func saddle(x, y float64) float64 {
	return math.Pow(y/xyrange*2, 2) - math.Pow(x/xyrange*2, 2)
}

func moguls(x, y float64) float64 {
	return math.Pow(math.Sin(x/xyrange*3*math.Pi), 2) * math.Cos(y/xyrange*3*math.Pi)
}

func eggBox(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 10
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	w.Header().Set("Content-Type", "image/svg+xml")

	var err error
	height, err = strconv.Atoi(r.Form.Get("height"))
	if err != nil {
		fmt.Printf("Error parsing height: %v", err)
	}

	if len(r.Form.Get("height")) > 0 {
		width, err = strconv.Atoi(r.Form.Get("height"))
	}

	if len(r.Form.Get("width")) > 0 {
		width, err = strconv.Atoi(r.Form.Get("width"))
	}

	peak_color_str := r.Form.Get("peak_color")
	if len(peak_color_str) > 0 {
		peak_color = peak_color_str
	}

	valley_color_str := r.Form.Get("valley_color")
	if len(valley_color_str) > 0 {
		valley_color = valley_color_str
	}

	writeSvg(w)
}
