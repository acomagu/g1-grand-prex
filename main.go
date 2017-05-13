package main

import (
	"fmt"
	"os"
)

// The number of cell by one side.
const l = 15

// State is whether the cell is black, white or empty.
type State int

const (
	Empty State = iota
	Me
	Opponent
)

// Field is whole game board.
type Field [l][l]State

func (field Field) copy() Field {
	var newField Field
	for y := 0; y < l; y++ {
		for x := 0; x < l; x++ {
			newField[y][x] = field[y][x]
		}
	}
	return newField
}

var verbose = false

func main() {
	if len(os.Args) >= 2 && os.Args[1] == "-v" {
		verbose = true
	}

	field := Field{}

	for {
		printField(field)

		var opponentY, opponentX int
		if n, err := fmt.Scanf("%d %d", &opponentY, &opponentX); n != 2 || err != nil {
			fmt.Printf("n: %s, error: %s", n, err)
			break
		}

		if opponentY != -1 && opponentX != -1 {
			field[opponentY][opponentX] = Opponent
		}

		printField(field)

		myY, myX, _ := calcNextPlacing(field)

		fmt.Printf("%d %d\n", myY, myX)
		field[myY][myX] = Me
	}
}

func printField(field Field) {
	if !verbose {
		return
	}

	fmt.Fprint(os.Stderr, "  ")
	
	for x := 0; x < l; x++ {
		fmt.Fprintf(os.Stderr, "%2d", x)
	}
	fmt.Fprintln(os.Stderr, )

	for y := 0; y < l; y++ {
		fmt.Fprintf(os.Stderr, "%2d", y)
		for x := 0; x < l; x++ {
			switch field[y][x] {
			case Empty:
				fmt.Fprint(os.Stderr, "  ")
			case Me:
				fmt.Fprint(os.Stderr, " o")
			case Opponent:
				fmt.Fprint(os.Stderr, " x")
			}
		}
		fmt.Fprintf(os.Stderr, "%2d\n", y)
	}

	fmt.Fprint(os.Stderr, "  ")
	for x := 0; x < l; x++ {
		fmt.Fprintf(os.Stderr, "%2d", x)
	}
	fmt.Fprintln(os.Stderr, )
}
