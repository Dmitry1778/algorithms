package main

import "fmt"

func QuickSort[T any](data []T, compare func(a, b T) bool) {
	if len(data) < 2 {
		return
	}
	left, right := 0, len(data)-1
	pivot := len(data) / 2
	data[pivot], data[right] = data[right], data[pivot]
	for i := range data {
		if compare(data[i], data[right]) {
			data[left], data[i] = data[i], data[left]
			left++
		}
	}
	data[left], data[right] = data[right], data[left]
	QuickSort(data[:left], compare)
	QuickSort(data[left+1:], compare)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 { // Базовый случай: если массив пустой или содержит один элемент, он уже отсортирован
		return arr
	}

	pivot := arr[len(arr)/2] // Выбор среднего элемента в качестве опорного
	var left, right []int

	for _, v := range arr {
		if v < pivot { // Сравнение с опорным элементом
			left = append(left, v) // Добавление элемента в левую часть
		} else if v > pivot {
			right = append(right, v) // Добавление элемента в правую часть
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...) // Рекурсивные вызовы и объединение результатов
}

func main() {
	slice := []int{2, 3, 1, 4}
	slice2 := []int{2, 3, 1, 4}
	QuickSort(slice, func(a, b int) bool { return a < b })
	fmt.Println(slice) // Выведет: [2 3 4 6 9 10]
	fmt.Println(quickSort(slice2))
}
