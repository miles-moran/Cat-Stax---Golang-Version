package points

import "fmt"

type Point struct {
	X int8
	Y int8
	Z int8
}

type PointArray struct {
	Points []Point
}

func (p Point) ListFormat() [3]int8 {
	listFormat := [3]int8{p.X, p.Y, p.Z}
	return listFormat
}

//Sorts series of points. Points with lowest Z values are placed first, then points with lowest Y values, then points with lowest X values
func Sort(P PointArray) PointArray {
	sorted := true
	for i := 0; i < len(P.Points)-1; i++ {
		p1 := P.Points[i]
		p2 := P.Points[i+1]
		if p1.Z > p2.Z {
			placeholder := p1
			P.Points[i] = p2
			P.Points[i+1] = placeholder
			sorted = false
		} else if p1.Z == p2.Z {
			if p1.Y > p2.Y {
				placeholder := p1
				P.Points[i] = p2
				P.Points[i+1] = placeholder
				sorted = false
			} else if p1.Y == p2.Y {
				if p1.X > p2.X {
					placeholder := p1
					P.Points[i] = p2
					P.Points[i+1] = placeholder
					sorted = false
				}
			}
		}
	}
	if sorted == false {
		fmt.Println("sorting")
		Sort(P)
	}
	return P
}
