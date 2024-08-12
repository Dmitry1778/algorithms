package main

import "fmt"

func main() {
	array := []int{1, 3, 2, 4, 6, 5}
	fmt.Println(selectionSort(array))
}
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	return arr
}
