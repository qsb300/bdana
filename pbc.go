package bdana

import (
	"math"
)

func PBC1d(x, boxl float64) float64 {
	_, xf := math.Modf(x / boxl)
	switch {
	case xf < -0.5:
		xf = xf + 1.0
	case xf > 0.5:
		xf = xf - 1.0
	}
	return boxl * xf
}

func PBC1ddst(x, boxl float64) float64 {
	_, xf := math.Modf(math.Abs(x) / boxl)
	if xf > 0.5 {
		xf = 1.0 - xf
	}
	return boxl * xf
}

func PBCDst2(x1, y1, z1, x2, y2, z2, boxl float64) float64 {
	xd := PBC1d(x1-x2, boxl)
	yd := PBC1d(y1-y2, boxl)
	zd := PBC1d(z1-z2, boxl)
	return xd*xd + yd*yd + zd*zd
}

func PBCDst2Multi(xyz [][3]float64, x, y, z, boxl float64) []float64 {
	n := len(xyz)
	v := make([]float64, n)
	for i := 0; i < n; i++ {
		v[i] = PBCDst2(x, y, z, xyz[i][0], xyz[i][1], xyz[i][2], boxl)
	}
	return v
}

func PBCImage(xyz [][3]float64, xc, yc, zc, boxl float64) {
	n := len(xyz)
	for i := 0; i < n; i++ {
		xyz[i][0] = PBC1d(xyz[i][0]-xc, boxl)
		xyz[i][1] = PBC1d(xyz[i][1]-yc, boxl)
		xyz[i][2] = PBC1d(xyz[i][2]-zc, boxl)
	}
}

func Min(v []float64) (m float64) {
	n := len(v)
	if n > 0 {
		m = v[0]
	}
	for i := 1; i < n; i++ {
		if v[i] < m {
			m = v[i]
		}
	}
	return
}
