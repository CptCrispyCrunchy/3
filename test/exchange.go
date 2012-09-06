package main

import (
	"nimble-cube/core"
	"nimble-cube/dump"
	"nimble-cube/mag"
)

func main() {
	N0, N1, N2 := 1, 4, 8
	size := [3]int{N0, N1, N2}
	cellsize := [3]float64{1e-9, 1e-9, 1e-9}

	m := core.MakeVectors(size)
	h := core.MakeVectors(size)

	mag.SetRegion(m, 0, 0,0, 1,N1, N2/2, mag.Uniform(0, 1, 0))
	mag.SetRegion(m, 0, 0,0, 1,N1, N2/2, mag.Uniform(0, 0, 1))

	Aex := 13e-12
	mag.Exchange6(m, h, cellsize, Aex)
	dump.Quick("h_ex", h[:])
} 
