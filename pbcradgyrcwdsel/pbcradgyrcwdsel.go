package main

import (
	"bdana"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func Read(filename string, rcut float64, showmodel bool) {
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
	rcut2 := rcut * rcut
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
			xc, yc, zc := bdana.Center(xyzpro)
			bdana.PBCImage(xyzpro, xc, yc, zc, boxl)
			bdana.PBCImage(xyzcwd, xc, yc, zc, boxl)
			cnt = cnt + 1
			rgpro := math.NaN()
			rgcwd := math.NaN()
			if !errpro {
				rgpro = bdana.RadGyr(xyzpro)
			}
			rgcwd = bdana.RadGyr(xyzcwd)
			fmt.Println("REMARK", cnt, npro, rgpro, ncwd, rgcwd, errcwd)
			if !errcwd && cnt > 1002 && showmodel && cnt%10 == 2 {
				fmt.Println("MODEL", cnt, rgpro, rgcwd)
				bdana.PDBWrite(xyzpro, 0, 0, " CA ", "pro")
				bdana.PDBWrite(xyzcwd, npro, npro, " CA ", "cwd")
				fmt.Println("ENDMDL")
			}
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
	rcut, _ := strconv.ParseFloat(os.Args[2], 64)
	showmodel, _ := strconv.ParseBool(os.Args[3])
	Read(os.Args[1], rcut, showmodel)
}
