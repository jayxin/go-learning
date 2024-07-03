package main

import "fmt"

func main() {
	a := 1
	if a < 5 {
		fmt.Println("a < 5")
	} else {
		fmt.Println("a >= 5")
	}

	i := 2
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4, 5 ,6")
	default:
		fmt.Println("Default")
	}

	Num := 2
	switch {
	case 0 <= Num && Num <= 3:
		fmt.Println("0-3")
	case 4 <= Num && Num <= 6:
		fmt.Println("4-6")
	case 7 <= Num && Num <= 9:
		fmt.Println("7-9")
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	sum = 0
	for {
		sum++
		if sum > 100 {
			break
		}
	}

	arr := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr)

JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}
	}

  myfunc()
}

func myfunc() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 5 {
		goto HERE
	}
}
