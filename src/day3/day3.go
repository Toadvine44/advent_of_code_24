package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func buildWordMap(matches [][][]byte, mulIndices [][]int, doIndices [][]int, dontIndices [][]int) map[int]string {

	wordMap := make(map[int]string)

	for i, mul := range matches { // add multiplies
		idx := mulIndices[i][0]
		wordMap[idx] = string(mul[0])
	}

	for _, doIdx := range doIndices { // add do's
		wordMap[doIdx[0]] = "do()"
	}

	for _, dontIdx := range dontIndices { // add don'ts
		wordMap[dontIdx[0]] = "dont()"
	}

	return wordMap
}

func evaluate(matches [][][]byte, mulIndices [][]int, doIndices [][]int, dontIndices [][]int) int {

	sum := 0

	// create a mapping from idx => string
	wordMap := buildWordMap(matches, mulIndices, doIndices, dontIndices)

	// now for the fun part
	do := true
	var keys []int
	for key := range wordMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		val := wordMap[key]

		if val == "do()" {
			do = true
		} else if val == "dont()" {
			do = false
		} else {
			if do {
				re := regexp.MustCompile(`\b\d{1,3}\b`)
				nums := re.FindAllString(val, 2)
				n1, _ := strconv.Atoi(nums[0])
				n2, _ := strconv.Atoi(nums[1])
				sum += n1 * n2
			}
		}
	}

	return sum
}

func uncorruptMemory(line string) int {

	mulPattern := `mul\((\d{1,3}),(\d{1,3})\)`
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`

	reMul := regexp.MustCompile(mulPattern)
	reDo := regexp.MustCompile(doPattern)
	reDont := regexp.MustCompile(dontPattern)

	mulVals := reMul.FindAllSubmatch([]byte(line), -1)
	mulIndices := reMul.FindAllIndex([]byte(line), -1)
	doIndices := reDo.FindAllIndex([]byte(line), -1)
	dontIndices := reDont.FindAllIndex([]byte(line), -1)

	return evaluate(mulVals, mulIndices, doIndices, dontIndices)
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ERROR OPENING FILE")
		os.Exit(2)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var inputStr string
	for scanner.Scan() {
		inputStr += scanner.Text()
	}
	fmt.Println("RESULT:", uncorruptMemory(inputStr))
}
