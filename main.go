package main

import (
	"fmt"

	"github.com/miles-moran/catStax/points"
)

func main() {

	point0 := points.Point{
		X: 2,
		Y: 2,
		Z: 0,
	}
	point1 := points.Point{
		X: 2,
		Y: 3,
		Z: 0,
	}

	var p = points.PointArray{}
	p.Points = []points.Point{point0, point1}

	for _, r := range points.GenerateRotations(p) {
		fmt.Println(r)
	}

}
