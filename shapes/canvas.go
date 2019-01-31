package shapes

import (
	"fmt"
	"strconv"
)

type Canvas struct {
	Points PointArray
}

func CanvasConsoleOutput(canvas Canvas) {
	maxZ := int(FindMax(canvas.Points, 2).Z) + 1
	maxY := int(FindMax(canvas.Points, 1).Y) + 1
	maxX := int(FindMax(canvas.Points, 0).X) + 1

	for z := 0; z < maxZ; z++ {
		if z != 0 {
			fmt.Println("")
		}
		for y := 0; y < maxY; y++ {
			fmt.Println("")
			for x := 0; x < maxX; x++ {
				occupant := "-"
				for _, p := range canvas.Points.Points {
					compare := [3]int8{
						int8(x),
						int8(y),
						int8(z),
					}
					if p.ListFormat() == compare {
						occupant = strconv.Itoa(p.occupant)
					} else {
					}
				}
				fmt.Print(occupant)
			}
		}
	}
	fmt.Printf("\n\n")
}
