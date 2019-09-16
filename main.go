package main

import (
	"fmt"
	"reflect"
)

func quicksort(arr []int, left int, right int) []int {
	if left < right {
		index := partition(arr, left, right)

		quicksort(arr, left, index-1)
		quicksort(arr, index+1, right)
	}
	return arr
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			i++
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}
	i++
	temp := arr[i]
	arr[i] = arr[right]
	arr[right] = temp

	return i
}

var testData = []struct {
	input, expectedOutput []int
	left, right           int
}{
	{[]int{}, []int{}, 0, 0},
	{[]int{1}, []int{1}, 0, 0},
	{[]int{1, 1}, []int{1, 1}, 0, 1},
	{[]int{1, 2, 1}, []int{1, 1, 2}, 0, 2},
	{[]int{13, 8, 5, 1, 1, 2, 3}, []int{1, 1, 2, 3, 5, 8, 13}, 0, 6},
}

func main() {
	for _, testCase := range testData {
		actual := quicksort(testCase.input, testCase.left, testCase.right)
		expected := testCase.expectedOutput

		if !reflect.DeepEqual(actual, expected) {
			fmt.Printf("Test failed: %v != %v\n", actual, expected)
		} else {
			fmt.Printf("Quicksorted: %v \n", actual)
		}
	}
}
