package bdana

import (
	"fmt"
	"strconv"
	"strings"
)

func GetXYZ(line string) (x, y, z float64, err error) {
	x, err = strconv.ParseFloat(strings.TrimSpace(line[30:38]), 64)
	if err != nil {
		return
	}
	y, err = strconv.ParseFloat(strings.TrimSpace(line[38:46]), 64)
	if err != nil {
		return
	}
	z, err = strconv.ParseFloat(strings.TrimSpace(line[46:54]), 64)
	if err != nil {
		return
	}
	return
}

func PDBLine(iatm, ires int, natm, nres string, x, y, z float64) (line string) {
	line = fmt.Sprintf("ATOM  %5d %4s %3s  %4d    %8.3f%8.3f%8.3f", iatm, natm, nres, ires, x, y, z)
	return
}

func PDBWrite(xyz [][3]float64, iatm, ires int, natm, nres string) {
	n := len(xyz)
	for i := 0; i < n; i++ {
		iatm++
		ires++
		line := PDBLine(iatm, ires, natm, nres, xyz[i][0], xyz[i][1], xyz[i][2])
		fmt.Println(line)
	}
}
