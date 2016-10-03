package bdana

import (
	"math"
)

func Center(xyz [][3]float64) (xc, yc, zc float64) {
	xc = 0.0
	yc = 0.0
	zc = 0.0
	n := len(xyz)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		xc += xyz[i][0]
		yc += xyz[i][1]
		zc += xyz[i][2]
	}
	xc = xc / float64(n)
	yc = yc / float64(n)
	zc = zc / float64(n)
	return
}

func radgyr2cen(xyz [][3]float64, xc, yc, zc float64) (rg float64) {
	n := len(xyz)
	if n == 0 {
		return 0.0
	}
	rg2 := 0.0
	for i := 0; i < n; i++ {
		rg2 += (xyz[i][0] - xc) * (xyz[i][0] - xc)
		rg2 += (xyz[i][1] - yc) * (xyz[i][1] - yc)
		rg2 += (xyz[i][2] - zc) * (xyz[i][2] - zc)
	}
	rg = math.Sqrt(rg2 / float64(n))
	return
}

func RadGyr(xyz [][3]float64) (rg float64) {
	n := len(xyz)
	if n == 0 {
		return 0.0
	}
	xc, yc, zc := Center(xyz)
	rg = radgyr2cen(xyz, xc, yc, zc)
	return
}
