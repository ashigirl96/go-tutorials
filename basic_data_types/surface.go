package main

import (
	"fmt"
	"math"
	"log"
	"net/http"
	"sync"
)

const (
	width, height = 600 * 3, 320 * 3
	cells = 100  // num of grid cells
	xyrange = 30.0  // axis range
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
	angle = math.Pi / 6  // angle of x, y axes (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)  // sin(30), cos(30)
var mu sync.Mutex
var count int


func main() {
	var svg string
	svg += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width: 0.7' " +
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := concer(i+1, j)
			bx, by := concer(i, j)
			cx, cy := concer(i, j+1)
			dx, dy := concer(i+1, j+1)
			svg += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	svg += fmt.Sprintf("</svg>")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintf(w, "%s", svg)
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}


func concer(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2-D SVG canvas
	sx := width/2 + (x-y)*cos30*xyscale
	sy := width/2 + (x-y)*sin30*xyscale - z*zscale
	return sx, sy
}


func f(x, y float64) float64 {
	r := math.Hypot(x, y)  // distance from (0, 0)
	return math.Sin(r) / r
}


