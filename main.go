package main

import (
	"github.com/miles-moran/catStax/shapes"
)

func main() {
	canvas := shapes.Generate3x3Canvas()
	shape0 := shapes.Generate2x2Square()
	shape1 := shapes.Generate3x1Rectangle()
	shape2 := shapes.Generate2x1Rectangle()
	shapesList := []shapes.Shape{
		shape0,
		shape1,
		shape2,
	}
	shapes.Solve(canvas, shapesList)
}
