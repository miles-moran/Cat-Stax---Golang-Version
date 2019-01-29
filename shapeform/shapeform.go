package shapeform

import (
	"github.com/miles-moran/catStax/points"
)

//From an array of points, this function generates the 48 ways

func GenerateRotations(P points.PointArray) []points.PointArray {
	rotations := []points.PointArray{}
	signs := []int8{-1, 1}
	for a := 0; a < 3; a++ {
		for _, signA := range signs {
			for b := 0; b < 3; b++ {
				for _, signB := range signs {
					for c := 0; c < 3; c++ {
						if a != b && b != c && c != a {
							for _, signC := range signs {
								rotation := points.PointArray{}
								for _, p := range P.Points {
									x := p.ListFormat()[a] * signA
									y := p.ListFormat()[b] * signB
									z := p.ListFormat()[c] * signC
									newPoint := points.Point{x, y, z}
									rotation.Points = append(rotation.Points, newPoint)
								}
								rotations = append(rotations, rotation)
							}
						}
					}
				}
			}
		}
	}
	return rotations
}

//func PositiveRotations(rotations []points.PointArray) []points.PointArray {

//}
