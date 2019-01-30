package shapes

import (
	"fmt"
)

func Solve(canvas Canvas, shapes []Shape) Canvas {

	locations := [][]PointArray{}
	for _, s := range shapes {
		location := LocationsGenerator(canvas, s)
		locations = append(locations, location)
	}
	shapeCount := len(shapes)

	locations0 := locations[0]
	for _, location0 := range locations0 {
		allPositions := []PointArray{
			location0,
		}
		if collisionsTest(allPositions) {
			continue
		} else {
			if shapeCount == 1 {
				claim(canvas, allPositions)
				fmt.Println("SOLUTION FOUND")
				return canvas
			} else {
				locations1 := locations[1]
				for _, location1 := range locations1 {
					allPositions := []PointArray{
						location0,
						location1,
					}
					if collisionsTest(allPositions) {
						continue
					} else {
						if shapeCount == 2 {
							claim(canvas, allPositions)
							fmt.Println("SOLUTION FOUND")
							return canvas
						} else {
							locations2 := locations[2]
							for _, location2 := range locations2 {
								allPositions := []PointArray{
									location0,
									location1,
									location2,
								}
								if collisionsTest(allPositions) {
									continue
								} else {
									if shapeCount == 3 {
										claim(canvas, allPositions)
										fmt.Println("SOLUTION FOUND")
										return canvas
									} else {
										locations3 := locations[3]
										for _, location3 := range locations3 {
											allPositions := []PointArray{
												location0,
												location1,
												location2,
												location3,
											}
											if collisionsTest(allPositions) {
												continue
											} else {
												if shapeCount == 4 {
													claim(canvas, allPositions)
													fmt.Println("SOLUTION FOUND")
													return canvas
												} else {
													locations4 := locations[4]
													for _, location4 := range locations4 {
														allPositions := []PointArray{
															location0,
															location1,
															location2,
															location3,
															location4,
														}
														if collisionsTest(allPositions) {
															continue
														} else {
															if shapeCount == 5 {
																claim(canvas, allPositions)
																fmt.Println("SOLUTION FOUND")
																return canvas
															} else {

															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return canvas
}

func collisionsTest(allPositions []PointArray) bool {
	for i, shapePositionI := range allPositions {
		for f, shapePositionF := range allPositions {
			if i != f {
				for _, pI := range shapePositionI.Points {
					for _, pF := range shapePositionF.Points {
						if pI.ListFormat() == pF.ListFormat() {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

func claim(canvas Canvas, allPositions []PointArray) Canvas {
	for c, cP := range canvas.Points.Points {
		for s, shape := range allPositions {
			for _, sP := range shape.Points {
				if cP.ListFormat() == sP.ListFormat() {
					canvas.Points.Points[c].Occupant = s
				} else {

				}
			}
		}
	}
	CanvasConsoleOutput(canvas)
	return canvas
}
