package points

type Point struct {
	X int8
	Y int8
	Z int8
}

type PointArray struct {
	Points []Point
}

//Returns an array of size 3. Each item in the array is x, y, and z. In that order. Used for rotation generation.
func (p Point) ListFormat() [3]int8 {
	listFormat := [3]int8{p.X, p.Y, p.Z}
	return listFormat
}

//Golang won't let me compare slices
//This function takes every XYZ from every point in a Point Array/Slice and turns it into a giant string which is used to compare
func (P PointArray) CompareFormat() string {
	compare := ""
	for _, p := range P.Points {
		for xyz := 0; xyz < 3; xyz++ {
			add := p.ListFormat()[xyz]
			compare += string(add)
		}
	}
	return compare
}

//Generates rotations based off of a point array
//The number of rotations varies. A symetrical 3d object like a cube is only going to have 1 rotation
//A 2 by 1 shape will generate 3

//Rotations also vary on if you are trying to solve a 2d puzzle or a 3d puzzle

func GenerateRotations(P PointArray) []PointArray {
	rotations := []PointArray{}
	signs := []int8{-1, 1}
	for a := 0; a < 3; a++ {
		for _, signA := range signs {
			for b := 0; b < 3; b++ {
				for _, signB := range signs {
					for c := 0; c < 3; c++ {
						if a != b && b != c && c != a {
							for _, signC := range signs {
								rotation := PointArray{}
								for _, p := range P.Points {
									x := p.ListFormat()[a] * signA
									y := p.ListFormat()[b] * signB
									z := p.ListFormat()[c] * signC
									newPoint := Point{x, y, z}
									rotation.Points = append(rotation.Points, newPoint)
								}
								rotation.Trim()                               //Moves user input to 0,0 on a coordinate plane
								rotation.FirstQuadrantConvert()               //Puts reflections/rotations of the point array within the first cuadrant
								Sort(rotation)                                //Sorts the rotation to easily check for repeats
								if repeatTest(rotations, rotation) == false { //checks to see of a rotation is not a duplicate
									rotations = append(rotations, rotation)
								}
							}
						}
					}
				}
			}
		}
	}
	return rotations
}

//Sort series of points. Points with lowest Z values are placed first, then points with lowest Y values, then points with lowest X values
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
		Sort(P)
	}
	return P
}

//takes a point array and a value 0-2. 0 for X, 1 for Y, 2 for Z. returns the point with the lowest value of X, Y, or Z
func FindMin(P *PointArray, xyz int) Point {
	min := P.Points[0]
	for _, p := range P.Points {
		if p.ListFormat()[xyz] < min.ListFormat()[xyz] {
			min = p
		}
	}
	return min
}

//Ensures that user input will hug the x y and z axis.
func (P *PointArray) Trim() *PointArray {
	for xyz := 0; xyz < 3; xyz++ {
		min := FindMin(P, xyz)
		if min.ListFormat()[xyz] != 0 {
			for i := range P.Points {
				if xyz == 0 {
					P.Points[i].X += min.X * -1
				} else if xyz == 1 {
					P.Points[i].Y += min.Y * -1
				} else {
					P.Points[i].Z += min.Z * -1
				}
			}
		}
	}
	return P
}

//When generating rotations of a PointArray, we get negative values. We need to to put all these points within the first quadrant.
func (P *PointArray) FirstQuadrantConvert() *PointArray {
	for xyz := 0; xyz < 3; xyz++ {
		min := FindMin(P, xyz)
		if min.ListFormat()[xyz] < 0 {
			for i := range P.Points {
				if xyz == 0 {
					P.Points[i].X += min.X * -1
				} else if xyz == 1 {
					P.Points[i].Y += min.Y * -1
				} else {
					P.Points[i].Z += min.Z * -1
				}
			}
		}
	}
	return P
}

//When generating rotations of symmetrical shapes, we get many copies of the same Point Arrays
//This returns a true or false if there is a repeat
func repeatTest(rotations []PointArray, rotation PointArray) bool {
	for _, r := range rotations {
		if rotation.CompareFormat() == r.CompareFormat() {
			return true
		}
	}
	return false
}
