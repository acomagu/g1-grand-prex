package main

import (
	"math/rand"
	"fmt"
	"time"
)

type Point struct {
	y, x int
}

var nilPoint = Point{-1, -1}

func calcNextPlacing(field Field) (int, int, error) {
	rand.Seed(time.Now().Unix())

	p := getDangerPoint(field, 0, Opponent, 1)
	if p != nilPoint {
		return p.y, p.x, nil
	}

	p = getDangerPoint(field, 0, Me, 3)
	if p != nilPoint {
		return p.y, p.x, nil
	}

	fmt.Println("RAND")
	for {
		y := rand.Intn(l)
		x := rand.Intn(l)
		if field[y][x] == Empty {
			return y, x, nil
		}
	}

	return 0, 0, nil
}

func getDangerPoint(field Field, n int, color State, depth int) (Point) {
	points := getNextPositions(field, color)

	// debugField := field.copy()
	// for _, p := range points {
	// 	debugField[p.y][p.x] = Me
	// }
	// printField(debugField)

	for _, p := range points {
		tmpField := field.copy()
		tmpField[p.y][p.x] = color
		if isFinish(tmpField, color) {
			return Point{p.y, p.x}
		}
	}
	if n < depth {
		for _, p := range points {
			tmpField := field.copy()
			tmpField[p.y][p.x] = color
			ps := getDangerPoint(tmpField, n+1, color, depth)
			if ps != nilPoint {
				return ps
			}
		}
	}
	return nilPoint
}

func isFinish(field Field, color State) bool {
	dys := []int{0, 1, 1, 1}
	dxs := []int{1, 1, 0, -1}

	for y := 0; y < l; y++ {
		for x := 0; x < l; x++ {
			for i, dy := range dys {
				dx := dxs[i]
				nChain := getNChain(field, y, x, dy, dx, color, 0)
				if nChain >= 5 {
					return true
				}
			}
		}
	}
	return false
}

func getNChain(field Field, y, x, dy, dx int, color State, n int) int {
	if !isValidPosition(y, x) || field[y][x] != color {
		return n
	}

	return getNChain(field, y+dy, x+dx, dy, dx, color, n+1)
}

func isValidPosition(y, x int) bool {
	return 0 <= y && y < l && 0 <= x && x < l
}

func getNextPositions(field Field, color State) ([]Point) {
	dys := []int{-1, 0, 1}
	dxs := []int{-1, 0, 1}

	pointsMap := make(map[Point]struct{})

	for y := 0; y < l; y++ {
		for x := 0; x < l; x++ {
			if field[y][x] == color {
				for _, dy := range dys {
					for _, dx := range dxs {
						if isValidPosition(y+dy, x+dx) && field[y+dy][x+dx] == Empty {
							pointsMap[Point{y+dy, x+dx}] = struct{}{}
						}
					}
				}
			}
		}
	}
	points := []Point{}
	for p, _ := range pointsMap {
		points = append(points, p)
	}
	return points
}
