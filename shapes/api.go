package shapes

func api() {
	canvas := Generate3x4Canvas()
	shape0 := Generate2x2L()
	shape1 := Generate3x2L()
	shape2 := Generate2x1Rectangle()
	shape3 := Generate1x1()
	shapesList := []Shape{
		shape0,
		shape1,
		shape2,
		shape3,
	}
	Solve(canvas, shapesList)
}
