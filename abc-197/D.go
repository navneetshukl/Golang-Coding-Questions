// https://atcoder.jp/contests/abc197/tasks/abc197_d

package main

import (
	"fmt"
	"math"
)

func D() {

	pi := math.Pi

	var n, x1, y1, x2, y2 float64
	fmt.Scan(&n, &x1, &y1, &x2, &y2)
	midx := (x1 + x2) / 2

	midy := (y1 + y2) / 2

	cos := math.Cos(2 * pi / n)
	sin := math.Sin(2 * pi / n)

	newX := midx + cos*(x1-midx) - sin*(y1-midy)
	newY := midy + cos*(y1-midy) + sin*(x1-midx)

	fmt.Printf("%.10f %.10f\n", newX, newY)
}
