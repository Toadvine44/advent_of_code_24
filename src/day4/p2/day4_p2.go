package main

import (
	"bufio"
	"fmt"
	"os"
)

func isMatch(xmasMtx [][]byte, i int, j int, char byte) bool {
	if i < 0 || j < 0 || i > len(xmasMtx)-1 || j > len(xmasMtx[i])-1 {
		return false
	}
	if xmasMtx[i][j] == char {
		return true
	}
	return false
}

func checkDiag1(xmasMtx [][]byte, i int, j int) bool {

	if isMatch(xmasMtx, i-1, j-1, 'M') && isMatch(xmasMtx, i+1, j+1, 'S') {
		return true
	}
	if isMatch(xmasMtx, i-1, j-1, 'S') && isMatch(xmasMtx, i+1, j+1, 'M') {
		return true
	}

	return false
}

func checkDiag2(xmasMtx [][]byte, i int, j int) bool {

	if isMatch(xmasMtx, i-1, j+1, 'M') && isMatch(xmasMtx, i+1, j-1, 'S') {
		return true
	}
	if isMatch(xmasMtx, i-1, j+1, 'S') && isMatch(xmasMtx, i+1, j-1, 'M') {
		return true
	}

	return false
}

func xmasSearch(xmasMtx [][]byte, i int, j int) bool {

	if checkDiag1(xmasMtx, i, j) && checkDiag2(xmasMtx, i, j) {
		return true
	}
	return false
}

func findXmas(xmasMtx [][]byte) int {

	cnt := 0

	// iterate through 2d slice, find all occurences of 'x'
	for i, row := range xmasMtx {
		for j, elem := range row {
			if elem == 'A' {
				res := xmasSearch(xmasMtx, i, j)
				if res {
					cnt++
				}
				// cnt += res
			}
		}
	}
	return cnt
}

func main() {

	file, err := os.Open("../input.txt")
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
