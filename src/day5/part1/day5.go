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

func validLine(ruleMap map[int][]int, line []int) (bool, error) {

	int2Idx := make(map[int]int)
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
						return false, nil
					}
				}
			}
		}
	}
	return true, nil
}

func calcPageNum(ruleMap map[int][]int, updates [][]int) (int, error) {

	sum := 0

	for _, line := range updates {
		isValidLine, err := validLine(ruleMap, line)
		if err != nil {
			return sum, err
		}
		if isValidLine {
			mid := len(line) / 2
			sum += line[mid]
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

	filename := "../example.txt"
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
