package day_03

import (
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

func (p Point) ManhattenDistance() int {
	x := p.X
	y := p.Y

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	return x + y
}

type LineSegment struct {
	Start, End Point
}

type Wire struct {
	Segments []LineSegment
}

func (w Wire) Points() (points map[Point]int) {
	points = make(map[Point]int)
	for _, line_segment := range w.Segments {
		for point := range line_segment.Points() {
			points[point] += 1
		}
	}

	return
}

func (l LineSegment) Points() (points map[Point]int) {
	points = make(map[Point]int)
	if l.Start.X == l.End.X { // Line is horizontal
		minY, maxY := min(l.Start.Y, l.End.Y), max(l.Start.Y, l.End.Y)
		for y := minY; y <= maxY; y++ {
			point := Point{l.Start.X, y}

			points[point] += 1
		}
	} else {
		minX, maxX := min(l.Start.X, l.End.X), max(l.Start.X, l.End.X)
		for x := minX; x <= maxX; x++ {
			point := Point{x, l.Start.Y}

			points[point] += 1
		}
	}

	return
}

func ParseWire(instructions string) (wire Wire) {
	inst := strings.Split(instructions, ",")
	x, y := 0, 0

	for _, instruction := range inst {
		op, numStr := instruction[0], instruction[1:]
		num, _ := strconv.Atoi(numStr)

		point := Point{x, y}

		switch op {
		case 'D':
			y -= num
			wire.Segments = append(wire.Segments, LineSegment{Start: point, End: Point{x, y}})
		case 'U':
			y += num
			wire.Segments = append(wire.Segments, LineSegment{Start: point, End: Point{x, y}})
		case 'L':
			x -= num
			wire.Segments = append(wire.Segments, LineSegment{Start: point, End: Point{x, y}})
		case 'R':
			x += num
			wire.Segments = append(wire.Segments, LineSegment{Start: point, End: Point{x, y}})
		}
	}
	return
}

func PartOne(input string) int {
	lines := strings.Split(input, "\n")
	wireOne := ParseWire(lines[0])
	wireTwo := ParseWire(lines[1])

	wOnePoints := wireOne.Points()
	wTwoPoints := wireTwo.Points()

	p := Point{}
	for point := range wOnePoints {
		if _, ok := wTwoPoints[point]; ok {
			if p.ManhattenDistance() == 0 || point.ManhattenDistance() < p.ManhattenDistance() {
				p = point
			}
		}
	}

	return p.ManhattenDistance()
}
