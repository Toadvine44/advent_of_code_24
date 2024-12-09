package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Check if arr is a valid record according to the
// following rules:
//   - must be either completely increasing or completely decreasing.
//   - difference d between any adjacent elements must be: 1 <= d <= 3
//
// Input:
//   - arr: an array of ints
//
// Returns:
//   - True if arr is valid, false otherwise
func isValidRecord(arr []int) bool {

	if len(arr) == 0 {
		return false
	}

	direction := 0
	for i, elem := range arr {
		if i == len(arr)-1 {
			break
		}

		next := arr[i+1]
		diff := elem - next
		if diff == 0 {
			return false
		}
		if math.Abs(float64(diff)) > 3 { // diff is too big, return false
			return false
		}
		if direction == 0 { // must be the first elem, set initial direction
			if diff < 0 {
				direction = -1
			} else {
				direction = 1
			}
		} else if direction == 1 { // positive direction
			if diff < 0 {
				return false
			}
		} else { // negative direction
			if diff > 0 {
				return false
			}
		}
	}
	return true
}

func hasAlternatives(record []int) bool {
	for i := range record {
		alt := make([]int, 0)
		alt = append(alt, record[:i]...)
		alt = append(alt, record[i+1:]...)
		if isValidRecord(alt) {
			return true
		}
	}
	return false
}

// Iterate through all records and aggregate the total
// amount of valid records.
//
// Input:
//   - records: a 2D array representing every record
//
// Returns:
//   - the number of valid records in records
//   - error in case of error, nil otherwise
func validateRecords(records [][]int) (int, error) {

	numValid := 0
	for _, record := range records {
		if isValidRecord(record) {
			numValid++
		} else {
			if hasAlternatives(record) {
				numValid++
			}
		}
	}
	return numValid, nil
}

// Scan the file line by line, and produce
// one slice of ints per line. Returns a two-dimensional
// slice.
//
// input:
//   - file: the file to read from
//
// Returns:
//   - 2D slice representing each slice for each line
//   - error in case of error, nil otherwise
func scanReport(file *os.File) ([][]int, error) {

	reports := [][]int{{}}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		reportSlice := []int{}
		fields := strings.Fields(line)
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("ERROR: line 27: error converting string to int")
				return reports, err
			}
			reportSlice = append(reportSlice, num)
		}
		reports = append(reports, reportSlice)
	}
	return reports, nil
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file")
		os.Exit(2)
	}
	defer file.Close()

	records, _ := scanReport(file)
	validRecords, _ := validateRecords(records)
	fmt.Println("VALID RECORDS:", validRecords)
}
