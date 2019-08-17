package main

import (
  "fmt"
  "reflect"
)

func quicksort(arr []int, left int, right int) []int {
  if left < right {
    index := partition(arr, left, right)

    quicksort(arr, left, index - 1)
    return quicksort(arr, index + 1, right)
  }
  return nil
}

func partition(arr []int, left int, right int) int {
  pivot := arr[right]
  i := left - 1

  for j := left; j < right; j++ {
    if arr[j] <= pivot {
      i++
    }
    temp := arr[i]
    arr[i] = arr[j]
    arr[j] =  temp
  }
  temp := arr[i + 1]
  arr[i + 1] = arr[right]
  arr[right] = temp

  return i + 1
}

var testData = []struct {
	input []int
	expectedOutput []int
  left, right int
}{
	{[]int{}, []int{}, 0, 0},
	{[]int{42}, []int{42}, 0, 0},
	{[]int{42, 23}, []int{23, 42}, 0, 1},
	{[]int{23, 42, 32, 64, 12, 4}, []int{4, 12, 23, 32, 42, 64}, 0, 5},
}

func main() {
  for _, testCase := range testData {
    actual := quicksort(testCase.input, testCase.left, testCase.right)
    expected := testCase.expectedOutput

    if !reflect.DeepEqual(actual, expected) {
      fmt.Printf("%v != %v\n", actual, expected)
    }
  }
}
