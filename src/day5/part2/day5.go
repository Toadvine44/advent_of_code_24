package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findMistake(ruleMap map[int][]int, line []int) (int, map[int]int, bool) {

	problemIndices := make(map[int]int)
	found := false
	int2Idx := make(map[int]int)
	var resultIdx int
	for i, num := range line {
		if _, ok := int2Idx[num]; !ok {
			int2Idx[num] = i
		}
	}

	for elemIdx, elem := range line {
		if prefixes, ok := ruleMap[elem]; ok {
			for _, prefix := range prefixes {
				if prefixIdx, ok := int2Idx[prefix]; ok {
					if !(prefixIdx < elemIdx) {
						resultIdx = elemIdx
						found = true
						problemIndices[line[prefixIdx]] = prefixIdx
					}
				}
			}
		}
		if found {
			break
		}
	}
	return resultIdx, problemIndices, found
}

func fixLine(ruleMap map[int][]int, line []int, initElemIdx int, initProblems map[int]int) []int {

	var found bool
	problemIndices := initProblems
	newLine := make([]int, len(line))

	copy(newLine, line)
	elemIdx := initElemIdx

	for {
		tempLine := make([]int, 0)
		preLine := make([]int, 0)
		postLine := make([]int, 0)

		for i, elem := range newLine {

			_, ok := problemIndices[elem]

			if i < elemIdx || ok { // this was either already before elem or needs to go before it
				preLine = append(preLine, elem)
			} else if i > elemIdx {
				postLine = append(postLine, elem)
			} else {
				continue
			}
		}

		// concat it all together
		tempLine = append(tempLine, preLine...)
		tempLine = append(tempLine, newLine[elemIdx])
		tempLine = append(tempLine, postLine...)

		copy(newLine, tempLine)
		elemIdx, problemIndices, found = findMistake(ruleMap, newLine)
		if !found {
			break
		}
	}

	return newLine
}

func calcPageNum(ruleMap map[int][]int, updates [][]int) (int, error) {

	sum := 0

	for _, line := range updates {
		elemIdx, problems, found := findMistake(ruleMap, line)
		if found {
			fixedLine := fixLine(ruleMap, line, elemIdx, problems)
			mid := len(fixedLine) / 2
			sum += fixedLine[mid]
		}
	}
	return sum, nil
}

func processInput(file *os.File) (map[int][]int, [][]int, error) {

	ruleMap := make(map[int][]int)
	updates := [][]int{}
	newLine := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			newLine = true
			continue
		}

		if !newLine { // rules
			numStrs := strings.Split(line, "|")
			prefix, prefixErr := strconv.Atoi(numStrs[0])
			key, keyErr := strconv.Atoi(numStrs[1])
			if keyErr != nil || prefixErr != nil {
				log.Println("ERROR PARSING RULES")
				return ruleMap, updates, errors.New("Error parsing rules string")
			}

			if prefixSlice, ok := ruleMap[key]; ok {
				ruleMap[key] = append(prefixSlice, prefix)
			} else {
				ruleMap[key] = []int{prefix}
			}
		} else { // updates
			updateStrSlice := strings.Split(line, ",")
			updateNums := []int{}
			for _, updateStr := range updateStrSlice {
				num, err := strconv.Atoi(updateStr)
				if err != nil {
					log.Println("ERROR PARSING UPDATE")
					return ruleMap, updates, errors.New("Error parsing rules string")
				}
				updateNums = append(updateNums, num)
			}
			updates = append(updates, updateNums)
		}
	}

	return ruleMap, updates, nil
}

func main() {

	filename := "../input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening file:", err)
		os.Exit(2)
	}

	ruleMap, updates, err := processInput(file)
	if err != nil {
		log.Println("ERROR returned from processInput")
		os.Exit(2)
	}
	result, err := calcPageNum(ruleMap, updates)
	if err != nil {
		log.Println("Error calculating sum:", err)
		os.Exit(2)
	}
	fmt.Println("RESULT:", result)
}
