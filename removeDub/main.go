package main

import "fmt"

func main() {

	nums := []int{1, 2, 2, 3, 1}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(removeDuplicatesMap(nums))
	// Работает с любым сравнимым типом
	fmt.Println(Distinct([]int{1, 2, 2, 3, 4, 4, 4, 5}))
	fmt.Println(Distinct([]string{"котик", "кошечка", "кот", "кошка"}))
}

func Distinct[T comparable](list []T) []T {
	var n bool
	fmt.Println(n)
	unique := make(map[T]bool)
	var result []T
	for _, item := range list {
		if !unique[item] {
			unique[item] = true
			result = append(result, item)
		}
	}
	return result
}

func removeDuplicatesMap(nums []int) int {
	k := 0
	result := []int{}
	encountered := map[int]bool{}

	for v := range nums {
		encountered[nums[v]] = true
		fmt.Println("map:", encountered)
	}

	for key, _ := range encountered {
		result = append(result, key)
		fmt.Println(result)
	}

	for i := 0; i < len(result); i++ {
		k++
	}
	return k
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	//result := []int{}
	result := []int{}
	k := 0
	currect := nums[0]
	for i := 1; i < len(nums); i++ {
		if currect != nums[i] {
			result = append(result, currect)
			k++
		} else if currect == nums[i] {
			result = append(result, currect)
		}

	}
	return k
}
