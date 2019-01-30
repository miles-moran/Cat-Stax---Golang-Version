package shapes

func LocationsGenerator(canvas Canvas, shape Shape) []PointArray {
	locations := []PointArray{}
	for _, r := range shape.Rotations {
		for _, l := range canvas.Points.Points {
			location := place(r, l, canvas)
			if len(location.Points) != 0 {
				locations = append(locations, location)
			}
		}
	}
	return locations
}

func place(rotation PointArray, location Point, canvas Canvas) PointArray {
	newLocation := PointArray{}
	for _, p := range rotation.Points {
		x := p.X + location.X - rotation.Points[0].X
		y := p.Y + location.Y - rotation.Points[0].Y
		z := p.Z + location.Z - rotation.Points[0].Z
		newPoint := Point{
			X: x,
			Y: y,
			Z: z,
		}
		if valid(newPoint, canvas) {
			newLocation.Points = append(newLocation.Points, newPoint)
		} else {
			newLocation = PointArray{}
			return newLocation
		}
	}
	return newLocation
}

func valid(p Point, canvas Canvas) bool {
	for _, c := range canvas.Points.Points {
		if c.ListFormat() == p.ListFormat() {
			return true
		}
	}
	return false
}
