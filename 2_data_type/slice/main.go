package main

import "fmt"

func main() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice []int = arr[:5]
	fmt.Println("Elements of arr: ")
	for _, v := range arr {
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of slice: ")
	for _, v := range slice {
		fmt.Print(v, " ")
	}

	slice = arr[5:]
	slice = arr[:]

	fmt.Println()

	s1 := make([]int, 5)

	s2 := make([]int, 5, 10)
	fmt.Println("Capacity: ", cap(s2))
	fmt.Println("Length: ", len(s2))
	s2 = append(s2, 1, 2, 3)

	s3 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1, s2, s3)

	for i := 0; i < len(s1); i++ {
		fmt.Println("s1[", i, "] = ", s1[i])
	}

	s4 := []int{8, 9, 10}
	s1 = append(s1, s4...)

	oldSlice := make([]int, 5, 10)
	newSlice := oldSlice[:3]
	newSlice = oldSlice[:7]
	fmt.Println(newSlice)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
  copy(slice2, slice1)
  fmt.Println(slice2) // [1 2 3]
}
