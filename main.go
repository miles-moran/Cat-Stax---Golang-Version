package main

import (
	"fmt"

	"github.com/miles-moran/catStax/points"
)

func main() {

	point0 := points.Point{
		X: 1,
		Y: 4,
		Z: 0,
	}
	point1 := points.Point{
		X: 0,
		Y: 4,
		Z: 0,
	}
	point2 := points.Point{
		X: 2,
		Y: 4,
		Z: 0,
	}
	point3 := points.Point{
		X: 0,
		Y: 0,
		Z: 5,
	}

	var p = points.PointArray{}
	p.Points = []points.Point{point0, point1, point2, point3}

	fmt.Println(p.Points)

	fmt.Println(points.Sort(p).Points)

}
