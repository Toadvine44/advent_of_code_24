package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	NORTH = iota
	EAST
	SOUTH
	WEST
)

type Position struct {
	i         int
	j         int
	direction int
}

type gridSquare struct {
	i int
	j int
}

func printFloor(floor [][]rune) {
	fmt.Println("########################## PRINTING FLOOR #######################")
	for _, row := range floor {
		fmt.Println(string(row))
	}
	fmt.Println("#################################################")
}

func parseInput(file *os.File) [][]rune {
	result := make([][]rune, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		result = append(result, line)
	}
	return result
}

func findGuard(floor [][]rune) Position {
	for i, row := range floor {
		for j := range row {
			if floor[i][j] == '^' {
				return Position{i, j, NORTH}
			}
			if floor[i][j] == '>' {
				return Position{i, j, EAST}
			}
			if floor[i][j] == 'V' {
				return Position{i, j, SOUTH}
			}
			if floor[i][j] == '<' {
				return Position{i, j, WEST}
			}
		}
	}
	return Position{-1, -1, -1}
}

func northMove(floor [][]rune, visited map[gridSquare]int, pos Position) (Position, [][]rune, map[gridSquare]int) {

	currGridSquare := gridSquare{pos.i, pos.j}
	visited[currGridSquare] = 1
	if pos.i-1 < 0 { // going out of map
		return Position{-1, -1, -1}, floor, visited
	}
	if floor[pos.i-1][pos.j] == '#' { // obstacle encountered, move right
		if pos.j+1 < len(floor[pos.i]) {
			newPos := Position{pos.i, pos.j + 1, EAST}
			floor[pos.i][pos.j] = '.'
			floor[pos.i][pos.j+1] = '>'
			return newPos, floor, visited
		} else {
			return Position{-1, -1, -1}, floor, visited
		}
	} else { // no obstacle, simply move ahead
		if pos.i-1 >= 0 {
			newPos := Position{pos.i - 1, pos.j, NORTH}
			floor[pos.i][pos.j] = '.'
			floor[pos.i-1][pos.j] = '^'
			return newPos, floor, visited
		} else {
			return Position{-1, -1, -1}, floor, visited
		}
	}
}

func eastMove(floor [][]rune, visited map[gridSquare]int, pos Position) (Position, [][]rune, map[gridSquare]int) {

	currGridSquare := gridSquare{pos.i, pos.j}
	visited[currGridSquare] = 1
	if pos.j+1 >= len(floor[pos.i]) { // going out of map
		return Position{-1, -1, -1}, floor, visited
	}

	if floor[pos.i][pos.j+1] == '#' { // obstacle encountered, move right
		if pos.i+1 < len(floor) {
			newPos := Position{pos.i + 1, pos.j, SOUTH}
			floor[pos.i][pos.j] = '.'
			floor[pos.i+1][pos.j] = 'V'
			return newPos, floor, visited
		} else {
			return Position{-1, -1, -1}, floor, visited
		}
	} else { // no obstacle, simply move ahead
		if pos.j+1 < len(floor[pos.i]) {
			newPos := Position{pos.i, pos.j + 1, EAST}
			floor[pos.i][pos.j] = '.'
			floor[pos.i][pos.j+1] = '>'
			return newPos, floor, visited
		} else {
			return Position{-1, -1, -1}, floor, visited
		}
	}

}

func southMove(floor [][]rune, visited map[gridSquare]int, pos Position) (Position, [][]rune, map[gridSquare]int) {
	currGridSquare := gridSquare{pos.i, pos.j}
	visited[currGridSquare] = 1
	if pos.i+1 >= len(floor) { // going out of map
		return Position{-1, -1, -1}, floor, visited
	}

	if floor[pos.i+1][pos.j] == '#' { // obstacle encountered, move right
		if pos.j-1 > 0 {
			newPos := Position{pos.i, pos.j - 1, WEST}
			floor[pos.i][pos.j] = '.'
			floor[pos.i][pos.j-1] = '<'
			return newPos, floor, visited
		} else {
			return Position{-1, -1, -1}, floor, visited
		}
	} else { // no obstacle, simply move ahead
		if pos.i+1 < len(floor) {
			newPos := Position{pos.i + 1, pos.j, SOUTH}
			floor[pos.i][pos.j] = '.'
			floor[pos.i+1][pos.j] = 'V'
			return newPos, floor, visited
		} else {
			return Position{-1, -1, -1}, floor, visited
		}
	}

}

func westMove(floor [][]rune, visited map[gridSquare]int, pos Position) (Position, [][]rune, map[gridSquare]int) {

	currGridSquare := gridSquare{pos.i, pos.j}
	visited[currGridSquare] = 1
	if pos.j-1 < 0 { // going out of map
		return Position{-1, -1, -1}, floor, visited
	}

	if floor[pos.i][pos.j-1] == '#' { // obstacle encountered, move right
		if pos.i-1 > 0 {
			newPos := Position{pos.i - 1, pos.j, NORTH}
			floor[pos.i][pos.j] = '.'
			floor[pos.i-1][pos.j] = '^'
			return newPos, floor, visited
		} else {
			return Position{-1, -1, -1}, floor, visited
		}
	} else { // no obstacle, simply move ahead
		if pos.j-1 >= 0 {
			newPos := Position{pos.i, pos.j - 1, WEST}
			floor[pos.i][pos.j] = '.'
			floor[pos.i][pos.j-1] = '<'
			return newPos, floor, visited
		} else {
			return Position{-1, -1, -1}, floor, visited
		}
	}

}

func makeMove(floor [][]rune, visited map[gridSquare]int, pos Position) (Position, [][]rune, map[gridSquare]int) {

	switch pos.direction {
	case NORTH:
		return northMove(floor, visited, pos)
	case EAST:
		return eastMove(floor, visited, pos)
	case SOUTH:
		return southMove(floor, visited, pos)
	case WEST:
		return westMove(floor, visited, pos)
	default:
		return Position{-1, -1, -1}, floor, visited
	}

	// return pos, floor, visited
}

func calcGuardPath(file *os.File) int {
	var pos Position
	floorMap := parseInput(file)
	visited := make(map[gridSquare]int)

	pos = findGuard(floorMap)
	for {
		pos, floorMap, visited = makeMove(floorMap, visited, pos)
		if pos.i == -1 {
			break
		}
	}
	return len(visited)
}

func main() {

	filename := "../input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR OPENING FILE:", filename)
		os.Exit(2)
	}

	fmt.Println("UNQIUE SQUARES VISITED:", calcGuardPath(file))
}
