package main

import (
	"fmt"

	"github.com/miles-moran/catStax/points"
)

func main() {

	point0 := points.Point{
		X: 0,
		Y: 0,
		Z: 0,
	}
	point1 := points.Point{
		X: 0,
		Y: 1,
		Z: 0,
	}
	point2 := points.Point{
		X: 1,
		Y: 0,
		Z: 0,
	}
	point3 := points.Point{
		X: 1,
		Y: 1,
		Z: 0,
	}
	point4 := points.Point{
		X: 0,
		Y: 0,
		Z: 1,
	}
	point5 := points.Point{
		X: 0,
		Y: 1,
		Z: 1,
	}
	point6 := points.Point{
		X: 1,
		Y: 0,
		Z: 1,
	}
	point7 := points.Point{
		X: 1,
		Y: 1,
		Z: 1,
	}

	var p = points.PointArray{}
	p.Points = []points.Point{point0, point1, point2, point3, point4, point5, point6, point7}

	for _, r := range points.GenerateRotations(p) {
		fmt.Println(r)
	}

}
