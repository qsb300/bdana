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
	var xyz [][3]float64
	for err == nil {
		if line[:3] == "REM" || line[:3] == "MOD" {
			natom = 0
			errf = false
			xyz = make([][3]float64, 0)
		} else if line[:3] == "END" {
			cnt = cnt + 1
			rg := math.NaN()
			if !errf {
				rg = bdana.RadGyr(xyz)
			}
			fmt.Println(cnt, natom, rg)
		} else {
			x, y, z, errl := bdana.GetXYZ(line)
			if errl == nil {
				xyz = append(xyz, [3]float64{x, y, z})
				natom = natom + 1
			} else {
				errf = true
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
