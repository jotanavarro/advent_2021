package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

type Diagram struct {
	width, height int
	grid          [][]int
}

// calculateDangerousPoints will return the number of points in a diagram that have a value of 2 or higher.
func (d *Diagram) calculateDangerousPoints() (result int) {
	for x := 0; x < d.width; x++ {
		for y := 0; y < d.height; y++ {
			if d.grid[x][y] > 1 {
				result++
			}
		}
	}
	return result
}

// drawNonDiagonalLine will draw a line between two points as long as it is horizontal or vertical in the diagram.
func (d *Diagram) drawNonDiagonalLine(origin Point, destination Point) {
	if origin.x == destination.x {
		start, end := origin.y, destination.y
		if origin.y > destination.y {
			// In case the line is drawn "backwards", we swap the start and end coordinates.
			start, end = end, start
		}
		for i := start; i <= end; i++ {
			d.grid[origin.x][i]++
		}
	} else if origin.y == destination.y {
		start, end := origin.x, destination.x
		if origin.x > destination.x {
			// In case the line is drawn "backwards", we swap the start and end coordinates.
			start, end = end, start
		}
		for i := start; i <= end; i++ {
			d.grid[i][origin.y]++
		}
	}
}

// drawLine draws a line which can be horizontal, vertical or at a 45-degree angle in the diagram.  Lines which are at
// a different angle will be ignored.
func (d *Diagram) drawLine(origin Point, destination Point, drawDiagonals bool) {
	d.resizeBoard(origin, destination)
	d.drawNonDiagonalLine(origin, destination)

	if drawDiagonals {
		degreeIsCorrect, angleInDegrees := d.diagonalAngle(origin, destination)
		// We only want points that are at 45, -45, 135 or -135 degrees.
		if degreeIsCorrect {
			// Since we are moving in 45-degrees, the distance on X and Y axis grow equally.  Because of this we can use
			// the absolute distance between X or Y coordinate to know the distance between our points.
			distance := int(math.Abs(float64(origin.x - destination.x)))
			var signX, signY int
			switch angleInDegrees {
			case -45:
				// To move at a -45-degree angleInDegrees, we increase X and decrease Y.
				signX, signY = 1, -1
			case 45:
				// To move at a 45-degree angleInDegrees, we increase X and Y.
				signX, signY = 1, 1
			case 135:
				// To move at a 135-degree angleInDegrees, we decrease X and increase Y.
				signX, signY = -1, 1
			case -135:
				// To move at a -135-degree angleInDegrees, we decrease X and Y.
				signX, signY = -1, -1
			}
			for i := 0; i <= distance; i++ {
				d.grid[origin.x+(i*signX)][origin.y+(i*signY)]++
			}
		}
	}
}

// resizeBoard will check, given an origin and destination Point, if the diagram where we want to draw them requires
// to be resized.  If so, it will take care of it.
func (d *Diagram) resizeBoard(origin Point, destination Point) {
	// We check the size of our board and resize if required.
	// first horizontally
	if origin.x+1 >= d.width || destination.x+1 >= d.width {
		if origin.x > destination.x {
			// extend extra length of origin
			d.extendWidth(origin)
		} else {
			// extend extra length of destination
			d.extendWidth(destination)
		}
	}
	// then vertically
	if origin.y+1 >= d.height || destination.y+1 >= d.height {
		if origin.y > destination.y {
			// extend extra length of origin
			d.extendHeight(origin)
		} else {
			// extend extra length of destination
			d.extendHeight(destination)
		}
	}
}

// extendWidth will extend horizontally the diagram we use, if the point has a larger X coordinate that the current
// board width, then we re-set it to its new value.
func (d *Diagram) extendWidth(point Point) {
	delta := point.x - d.width + 1

	if delta > 0 {
		newSection := make([][]int, delta)
		for i := range newSection {
			newSection[i] = make([]int, d.height)
		}

		d.grid = append(d.grid, newSection...)
		d.width = point.x + 1
	}
}

// extendHeight will extend vertically the diagram we use, if the point has a larger Y coordinate that the current
// board height, then we re-set it to its new value.
func (d *Diagram) extendHeight(point Point) {
	delta := point.y - d.height + 1

	if delta > 0 {
		for i := range d.grid {
			d.grid[i] = append(d.grid[i], make([]int, delta)...)
		}
		d.height = point.y + 1
	}
}

func (d *Diagram) drawDiagram() {
	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			fmt.Printf("%d ", d.grid[x][y])
		}
		fmt.Printf("\n")
	}
}

// diagonalAngle will return true in case the two provided points are at a 45-degree angle.
func (d *Diagram) diagonalAngle(origin Point, destination Point) (bool, float64) {
	radianAngle := math.Atan2(float64(destination.y-origin.y), float64(destination.x-origin.x))
	degreeAngle := radianAngle * (180 / math.Pi)

	return degreeAngle == 45 || degreeAngle == -45 || degreeAngle == 135 || degreeAngle == -135, degreeAngle
}
