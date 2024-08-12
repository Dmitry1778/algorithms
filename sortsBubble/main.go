package main

import "fmt"

func main() {
	array := []int{1, 3, 2, 4, 6, 5}
	fmt.Println(bubble(array))
	array1 := []int{1, 3, 2, 4, 6, 5}
	fmt.Println(bubbleSM(array1))
	array2 := []int{1, 3, 2, 4, 6, 5}
	bubbleSort(array2)
}

func bubble(array []int) []int {
	var temp int
	for i := 0; i < len(array); i++ {
		for j := 0; j < (len(array) - i - 1); j++ {
			if array[j] > array[j+1] {
				temp = array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
	return array
}

func bubbleSM(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < (len(array) - i - 1); j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func bubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("\nAfter Bubble Sorting")
	for _, val := range arr {
		fmt.Println(val)
	}
}
