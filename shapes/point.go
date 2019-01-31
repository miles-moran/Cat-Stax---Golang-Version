package shapes

type Point struct {
	X        int8 `json:"x"`
	Y        int8 `json:"y"`
	Z        int8 `json:"z"`
	occupant int  `json:"occupant,omitempty"`
}

//func NewPoint() *Point {
//	point := &Point{}
//	point.Occupant = new(Point).Occupant
//	return point
//}

//Returns an array of size 3. Each item in the array is x, y, and z. In that order. Used for rotation generation.
func (p Point) ListFormat() [3]int8 {
	listFormat := [3]int8{p.X, p.Y, p.Z}
	return listFormat
}
