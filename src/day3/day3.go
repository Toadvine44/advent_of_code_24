package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func evaluate(matches [][][]byte) int {
	sum := 0

	for _, match := range matches {
		// this is going to be of pattern ['m','u','l','(',....,')']
		str := string(match[0])
		re := regexp.MustCompile(`\b\d{1,3}\b`)
		nums := re.FindAllString(str, 2)
		n1, _ := strconv.Atoi(nums[0])
		n2, _ := strconv.Atoi(nums[1])
		sum += n1 * n2
	}
	return sum
}

func uncorruptMemory(line string) int {

	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllSubmatch([]byte(line), -1)
	return evaluate(matches)
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ERROR OPENING FILE")
		os.Exit(2)
	}

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		product := uncorruptMemory(line)
		sum += product
	}
	fmt.Println("SUM OF ALL VALID MULTIPLICATIONS:", sum)
}
