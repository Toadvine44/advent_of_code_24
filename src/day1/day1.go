package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

// Compute distance scans the input file, and builds
// the two array inputs. Then, it sorts the arrays and
// returns the pairwise aggregated distance between them.
//
// Input:
//   - file: the input file
//
// Returns:
//   - Distance between the arrays
func sortInput(file *os.File) ([]int, []int) {

	scanner := bufio.NewScanner(file)
	arr1 := []int{}
	arr2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		var num1, num2 int
		_, err := fmt.Sscanf(line, "%d %d", &num1, &num2)
		if err != nil {
			fmt.Println("ERROR: error parsing line.")
		}

		arr1 = append(arr1, num1)
		arr2 = append(arr2, num2)
	}

	arr1 = quickSort(arr1, 0, len(arr1)-1)
	arr2 = quickSort(arr2, 0, len(arr2)-1)
	return arr1, arr2
}

// Computes the aggregated pairwise distance between two
// sorted int arrays of equal length. The algorithm computes
// the absolute value of the difference between elements at
// each index, then aggregates those differences.
//
// Input:n"
//   - arr1: an array of ints
//   - arr2: an array of ints
//
// Returns:
//   - the distance between arr1 and arr2, -1 in case of error.
func arrDistance(arr1 []int, arr2 []int) int {
	dist := 0
	if len(arr1) == len(arr2) {
		for i := 0; i < len(arr1); i++ {
			diff := int(math.Abs(float64(arr1[i] - arr2[i])))
			dist += diff
		}
		return dist
	} else {
		return -1
	}

}

// Implementation of the quick sort algorithm. If the array
// length is greater than 1, we will partition the array, and
// then recursively sort the two partitions.
//
// Input:
//   - arr: an array of integers
//   - low: the low bound idx for sorting
//   - high: the high bound idx for sorting
//
// Returns:
//   - the sorted array
func quickSort(arr []int, low int, high int) []int {
	if low < high {
		pivotIdx := partition(arr, low, high)
		quickSort(arr, low, pivotIdx-1)
		quickSort(arr, pivotIdx+1, high)
	}
	return arr
}

// Partitions arr by selecting a pivot (idx high), then
// placing any elems less than pivot to the left, and any
// elems greater than pivot to the right.
//
// Input:
//   - arr: an array of integers
//   - low: the low bound idx to partition
//   - high: the high bound idx to partition
//
// Returns:
//   - The index of the pivot after partition is completed
func partition(arr []int, low int, high int) int {
	i := low
	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			swapElem(arr, i, j)
			i++
		}
	}
	swapElem(arr, i, high)
	return i
}

// Helper function for swapping two ints in an
// integer array by index.
//
// Input:
//   - arr: array of integers
//   - i: an index
//   - j: an index
//
// Returns: None
func swapElem(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// Computes the similarity between arr1 and arr2.
func computeSimilarity(arr1 []int, arr2 []int) int {
	sim := 0
	simMap := make(map[int]int)
	for _, num := range arr2 { // Store the count for each num in arr2
		val, e := simMap[num]
		if e {
			simMap[num] = val + 1
		} else {
			simMap[num] = 1
		}
	}

	for _, num := range arr1 {
		val, e := simMap[num]
		if e { // similarity found
			sim += num * val
		}
	}

	return sim
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("ERROR OPENING INPUT FILE.\n")
	}
	defer file.Close()

	arr1, arr2 := sortInput(file)

	fmt.Println("::::: PART 1 ::::: Total computed distance:", arrDistance(arr1, arr2))
	fmt.Println("::::: PART 2 ::::: Similarity score:", computeSimilarity(arr1, arr2))
}
