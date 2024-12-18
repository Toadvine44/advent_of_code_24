package main

import (
	"bufio"
	"fmt"
	"os"
)

func print2dSlice(mtx [][]byte) {
	fmt.Println("================= PRINTING 2D Slice =================")
	fmt.Println("BYTE REPRESENTATION")
	for _, row := range mtx {
		fmt.Println(row)
	}
	fmt.Println("======================================================")
	fmt.Println("STRING REPRESENTATION")
	for _, row := range mtx {
		fmt.Println(string(row))
	}
	fmt.Println("=======================================================")
}

func isMatch(xmasMtx [][]byte, i int, j int, char byte) bool {
	if i < 0 || j < 0 || i > len(xmasMtx)-1 || j > len(xmasMtx[i])-1 {
		return false
	}
	if xmasMtx[i][j] == char {
		return true
	}
	return false
}

func searchEast(xmasMtx [][]byte, i int, j int) bool {
	return (isMatch(xmasMtx, i, j, 'X') &&
		isMatch(xmasMtx, i, j+1, 'M') &&
		isMatch(xmasMtx, i, j+2, 'A') &&
		isMatch(xmasMtx, i, j+3, 'S'))
}

func searchWest(xmasMtx [][]byte, i int, j int) bool {
	return (isMatch(xmasMtx, i, j, 'X') &&
		isMatch(xmasMtx, i, j-1, 'M') &&
		isMatch(xmasMtx, i, j-2, 'A') &&
		isMatch(xmasMtx, i, j-3, 'S'))
}

func searchNorth(xmasMtx [][]byte, i int, j int) bool {
	return (isMatch(xmasMtx, i, j, 'X') &&
		isMatch(xmasMtx, i-1, j, 'M') &&
		isMatch(xmasMtx, i-2, j, 'A') &&
		isMatch(xmasMtx, i-3, j, 'S'))
}

func searchSouth(xmasMtx [][]byte, i int, j int) bool {
	return (isMatch(xmasMtx, i, j, 'X') &&
		isMatch(xmasMtx, i+1, j, 'M') &&
		isMatch(xmasMtx, i+2, j, 'A') &&
		isMatch(xmasMtx, i+3, j, 'S'))
}

func searchNorthEast(xmasMtx [][]byte, i int, j int) bool {
	return (isMatch(xmasMtx, i, j, 'X') &&
		isMatch(xmasMtx, i-1, j+1, 'M') &&
		isMatch(xmasMtx, i-2, j+2, 'A') &&
		isMatch(xmasMtx, i-3, j+3, 'S'))
}

func searchNorthWest(xmasMtx [][]byte, i int, j int) bool {
	return (isMatch(xmasMtx, i, j, 'X') &&
		isMatch(xmasMtx, i-1, j-1, 'M') &&
		isMatch(xmasMtx, i-2, j-2, 'A') &&
		isMatch(xmasMtx, i-3, j-3, 'S'))
}

func searchSouthEast(xmasMtx [][]byte, i int, j int) bool {
	return (isMatch(xmasMtx, i, j, 'X') &&
		isMatch(xmasMtx, i+1, j+1, 'M') &&
		isMatch(xmasMtx, i+2, j+2, 'A') &&
		isMatch(xmasMtx, i+3, j+3, 'S'))
}

func searchSouthWest(xmasMtx [][]byte, i int, j int) bool {
	return (isMatch(xmasMtx, i, j, 'X') &&
		isMatch(xmasMtx, i+1, j-1, 'M') &&
		isMatch(xmasMtx, i+2, j-2, 'A') &&
		isMatch(xmasMtx, i+3, j-3, 'S'))
}

func xmasSearch(xmasMtx [][]byte, i int, j int) int {

	cnt := 0

	if searchEast(xmasMtx, i, j) {
		cnt++
	}
	if searchWest(xmasMtx, i, j) {
		cnt++
	}
	if searchNorth(xmasMtx, i, j) {
		cnt++
	}
	if searchSouth(xmasMtx, i, j) {
		cnt++
	}
	if searchNorthEast(xmasMtx, i, j) {
		cnt++
	}
	if searchNorthWest(xmasMtx, i, j) {
		cnt++
	}
	if searchSouthEast(xmasMtx, i, j) {
		cnt++
	}
	if searchSouthWest(xmasMtx, i, j) {
		cnt++
	}

	return cnt
}

func findXmas(xmasMtx [][]byte) int {

	cnt := 0

	// iterate through 2d slice, find all occurences of 'x'
	for i, row := range xmasMtx {
		for j, elem := range row {
			if elem == 'X' {
				res := xmasSearch(xmasMtx, i, j)
				cnt += res
			}
		}
	}
	return cnt
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ERROR OPENING FILE")
		os.Exit(2)
	}

	var xmasMtx [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		xmasMtx = append(xmasMtx, []byte(line))
	}

	result := findXmas(xmasMtx)
	fmt.Println("result: ", result)
}
