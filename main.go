package main

import (
	"github.com/miles-moran/catStax/shapes"
)

func main() {
	canvas := shapes.Generate3x4Canvas()
	shape0 := shapes.Generate2x2L()
	shape1 := shapes.Generate3x2L()
	shape2 := shapes.Generate2x1Rectangle()
	shape3 := shapes.Generate1x1()
	shapesList := []shapes.Shape{
		shape0,
		shape1,
		shape2,
		shape3,
	}
	shapes.Solve(canvas, shapesList)
}
