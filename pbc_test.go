package bdana

import (
	"math"
	"testing"
)

const TOLERANCE = 0.000001

func TestPBC1d(t *testing.T) {
	cases := []struct {
		in, boxl, want float64
	}{
		{1.1, 1.0, 0.1},
		{-1.7, 1.0, 0.3},
		{-5.2, 1.0, -0.2},
		{6.8, 1.0, -0.2},
		{-1.5, 1.0, -0.5},
		{100.5, 1.0, 0.5},
	}
	for _, c := range cases {
		got := PBC1d(c.in, c.boxl)
		if math.Abs(got-c.want) > TOLERANCE {
			t.Errorf("PBC1d(%f,%f) == %f, want %f", c.in, c.boxl, got, c.want)
		}
	}
}

func TestPBC1ddst(t *testing.T) {
	cases := []struct {
		in, boxl, want float64
	}{
		{1.1, 1.0, 0.1},
		{-1.7, 1.0, 0.3},
		{-5.2, 1.0, -0.2},
		{6.8, 1.0, -0.2},
		{-1.5, 1.0, -0.5},
		{100.5, 1.0, 0.5},
	}
	for _, c := range cases {
		got := PBC1ddst(c.in, c.boxl)
		if math.Abs(got-math.Abs(c.want)) > TOLERANCE {
			t.Errorf("PBC1d(%f,%f) == %f, want %f", c.in, c.boxl, got, math.Abs(c.want))
		}
	}
}

func TestMin(t *testing.T) {
	cases := []struct {
		in   []float64
		want float64
	}{
		{[]float64{1.0, 2.0, 3.0}, 1.0},
		{[]float64{2.0, 1.0, 3.0}, 1.0},
		{[]float64{1.0}, 1.0},
		{[]float64{}, 0.0},
	}
	for _, c := range cases {
		got := Min(c.in)
		if math.Abs(got-c.want) > TOLERANCE {
			t.Errorf("got %f, want %f", got, c.want)
		}
	}
}
