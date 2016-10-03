package main

import (
	"bdana"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func Read(filename string) {
	boxl := 510.0
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	rb := bufio.NewReader(f)
	line, err := rb.ReadString('\n')
	cnt := 0
	npro := 0
	ncwd := 0
	errpro := false
	errcwd := false
	rcut2 := (30.0 + 15.0) * (30.0 + 15.0)
	var xyzpro [][3]float64
	var xyzcwd [][3]float64
	for err == nil {
		if line[:3] == "REM" || line[:3] == "MOD" {
			npro = 0
			ncwd = 0
			errpro = false
			errcwd = false
			xyzpro = make([][3]float64, 0)
			xyzcwd = make([][3]float64, 0)
		} else if line[:3] == "END" {
			cnt = cnt + 1
			rgpro := math.NaN()
			rgcwd := math.NaN()
			if !errpro {
				rgpro = bdana.RadGyr(xyzpro)
			}
			rgcwd = bdana.RadGyr(xyzcwd)
			fmt.Println(cnt, npro, rgpro, ncwd, rgcwd, errcwd)
		} else {
			x, y, z, errl := bdana.GetXYZ(line)
			if errl == nil {
				if line[17:20] == "cwd" {
					dst2pro := bdana.PBCDst2Multi(xyzpro, x, y, z, boxl)
					dmin2 := bdana.Min(dst2pro)
					//fmt.Println("cnt",cnt,"dmin2",dmin2)
					if dmin2 < rcut2 {
						xyzcwd = append(xyzcwd, [3]float64{x, y, z})
						ncwd = ncwd + 1
					}
				} else {
					xyzpro = append(xyzpro, [3]float64{x, y, z})
					npro = npro + 1
				}
			} else {
				if line[17:20] == "cwd" {
					errcwd = true
				} else {
					errpro = true
				}
			}
		}
		line, err = rb.ReadString('\n')
	}
	//    fmt.Println(cnt)
	if err != io.EOF {
		fmt.Println(err)
	}
	return
}

func main() {
	Read(os.Args[1])
}
