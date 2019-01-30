package shapes

import (
	"fmt"
)

type Shape struct {
	Rotations []PointArray
}

//Prints rotations of a shape in console
func (s Shape) PrintRotations() {
	for _, r := range s.Rotations {
		fmt.Println(r)
	}
}
