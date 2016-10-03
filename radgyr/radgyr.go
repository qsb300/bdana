package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func Read(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	rb := bufio.NewReader(f)
	line, err := rb.ReadString('\n')
	cnt := 0
	natom := 0
	errf := false
	var x, y, z float64
	var xyz [][3]float64
	for err == nil {
		if line[:3] == "REM" || line[:3] == "MOD" {
			natom = 0
			xyz = make([][3]float64, 0)
		} else if line[:3] == "END" {
			cnt = cnt + 1
			rg := math.NaN()
			if !errf {
				rg = radgyr(xyz)
			}
			fmt.Println(cnt, natom, rg)
		} else {
			x, err = strconv.ParseFloat(strings.TrimSpace(line[30:38]), 64)
			if err != nil {
				errf = true
			}
			//fmt.Println(line[30:38],xc,err)
			y, err = strconv.ParseFloat(strings.TrimSpace(line[38:46]), 64)
			if err != nil {
				errf = true
			}
			//fmt.Println(line[38:46],xc,err)
			z, err = strconv.ParseFloat(strings.TrimSpace(line[46:54]), 64)
			if err != nil {
				errf = true
			}
			//fmt.Println(line[46:54],xc,err)
			xyz = append(xyz, [3]float64{x, y, z})
			natom = natom + 1
		}
		line, err = rb.ReadString('\n')
	}
	//    fmt.Println(cnt)
	if err != io.EOF {
		fmt.Println(err)
	}
	return
}

func radgyr(xyz [][3]float64) (rg float64) {
	n := len(xyz)
	xc := 0.0
	yc := 0.0
	zc := 0.0
	rg2 := 0.0
	for i := 0; i < n; i++ {
		xc += xyz[i][0]
		yc += xyz[i][1]
		zc += xyz[i][2]
	}
	xc = xc / float64(n)
	yc = yc / float64(n)
	zc = zc / float64(n)
	for i := 0; i < n; i++ {
		rg2 += (xyz[i][0] - xc) * (xyz[i][0] - xc)
		rg2 += (xyz[i][1] - yc) * (xyz[i][1] - yc)
		rg2 += (xyz[i][2] - zc) * (xyz[i][2] - zc)
	}
	rg = math.Sqrt(rg2 / float64(n))
	return
}

func main() {
	Read(os.Args[1])
}
