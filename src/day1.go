package main

import (
	"fmt"
	"math"
)

func main() {

	// TODO: add I/O so user can input two arrays on cmd line
	arr1 := []int{8, 12, 3, 0, 6}
	arr2 := []int{9, 4, 3, 0, 10}

	if !(len(arr1) == len(arr2)) {
		fmt.Println("ERROR: Please input arrays of equal length")
		return
	}

	fmt.Println("INPUT ARRAY1:", arr1)
	fmt.Println("INPUT ARRAY2:", arr2)

	arr1 = quickSort(arr1, 0, len(arr1)-1)
	arr2 = quickSort(arr2, 0, len(arr2)-1)

	fmt.Println("ARRAY1 SORTED:", arr1)
	fmt.Println("ARRAY2 SORTED:", arr2)

	dist := computeDistance(arr1, arr2)
	fmt.Println("Distance:", dist)
}

// Computes the aggregated pairwise distance between two
// sorted int arrays of equal length. The algorithm computes
// the absolute value of the difference between elements at
// each index, then aggregates those differences.
//
// Input:
//   - arr1: an array of ints
//   - arr2: an array of ints
//
// Returns:
//   - the distance between arr1 and arr2, -1 in case of error.
func computeDistance(arr1 []int, arr2 []int) int {
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
