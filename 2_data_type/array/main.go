package main

import "fmt"

func main() {
	var arr0 [32]int
	arr0[0] = 0

	var arr1 [3][5]int
	arr1[1][2] = 1

	var arr2 [2]([2]([2]float64)) // iff [2][2][2]float64
	arr2[0][1][1] = 2.0

	var arr3 [1000]*float64
	arr3[3] = nil

	fmt.Println(len(arr0))

	for i := 0; i < len(arr0); i++ {
		arr0[i] = i
	}

	for i, v := range arr0 {
		fmt.Println(i, v)
	}

	array := [5]int{1, 2, 3, 4, 5}
	modify(array)
	fmt.Println("In main(), array values: ", array)
}

func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify(), array values: ", array)
}
