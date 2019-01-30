package shapes

func Generate3x3Canvas() Canvas {
	point0 := Point{
		X: 0,
		Y: 0,
		Z: 0,
	}
	point1 := Point{
		X: 1,
		Y: 0,
		Z: 0,
	}
	point2 := Point{
		X: 2,
		Y: 0,
		Z: 0,
	}
	point3 := Point{
		X: 0,
		Y: 1,
		Z: 0,
	}
	point4 := Point{
		X: 1,
		Y: 1,
		Z: 0,
	}
	point5 := Point{
		X: 2,
		Y: 1,
		Z: 0,
	}
	point6 := Point{
		X: 0,
		Y: 2,
		Z: 0,
	}
	point7 := Point{
		X: 1,
		Y: 2,
		Z: 0,
	}
	point8 := Point{
		X: 2,
		Y: 2,
		Z: 0,
	}
	canvasPoints := PointArray{}
	canvasPoints.Points = []Point{point0, point1, point2, point3, point4, point5, point6, point7, point8}
	canvas := Canvas{
		Points: canvasPoints,
	}
	return canvas
}

func Generate3x1Rectangle() Shape {
	point0 := Point{
		X: 0,
		Y: 0,
		Z: 0,
	}
	point1 := Point{
		X: 1,
		Y: 0,
		Z: 0,
	}
	point2 := Point{
		X: 2,
		Y: 0,
		Z: 0,
	}
	canvasPoints := PointArray{}
	canvasPoints.Points = []Point{point0, point1, point2}
	canvas := Shape{
		Rotations: GenerateRotations(canvasPoints),
	}
	return canvas
}

func Generate2x1Rectangle() Shape {
	point0 := Point{
		X: 0,
		Y: 0,
		Z: 0,
	}
	point1 := Point{
		X: 1,
		Y: 0,
		Z: 0,
	}

	canvasPoints := PointArray{}
	canvasPoints.Points = []Point{point0, point1}
	canvas := Shape{
		Rotations: GenerateRotations(canvasPoints),
	}
	return canvas
}

func Generate2x2Square() Shape {
	point0 := Point{
		X: 0,
		Y: 0,
		Z: 0,
	}
	point1 := Point{
		X: 1,
		Y: 0,
		Z: 0,
	}
	point2 := Point{
		X: 0,
		Y: 1,
		Z: 0,
	}
	point3 := Point{
		X: 1,
		Y: 1,
		Z: 0,
	}

	canvasPoints := PointArray{}
	canvasPoints.Points = []Point{point0, point1, point2, point3}
	canvas := Shape{
		Rotations: GenerateRotations(canvasPoints),
	}
	return canvas
}
