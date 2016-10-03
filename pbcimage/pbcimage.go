package main

import (
	"bdana"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Read(filename string, cenpro bool) {
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
			var xc, yc, zc float64
			if cenpro {
				xc, yc, zc = bdana.Center(xyzpro)
			} else {
				xc, yc, zc = 0.0, 0.0, 0.0
			}
			bdana.PBCImage(xyzpro, xc, yc, zc, boxl)
			bdana.PBCImage(xyzcwd, xc, yc, zc, boxl)
			fmt.Println("REMARK", cnt, npro, ncwd)
			if !errpro && !errcwd && cnt == 1001 {
				fmt.Println("MODEL", cnt, npro, ncwd)
				bdana.PDBWrite(xyzpro, 0, 0, " CA ", "pro")
				bdana.PDBWrite(xyzcwd, npro, npro, " CA ", "cwd")
				fmt.Println("ENDMDL")
			}
		} else {
			x, y, z, errl := bdana.GetXYZ(line)
			if errl == nil {
				if line[17:20] == "cwd" {
					xyzcwd = append(xyzcwd, [3]float64{x, y, z})
					ncwd = ncwd + 1
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
	cenpro, err := strconv.ParseBool(os.Args[2])
	if err == nil {
		Read(os.Args[1], cenpro)
	}
}
