package main

import "fmt"

func main() {
	var str string
	str = "Hello world"
	ch := str[0]
	fmt.Printf("The length of \"%s\" is %d.\n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	str1 := "Hello"
	//str1[0] = 'A' // error!
	fmt.Println(str1)

	fmt.Println("Hello " + "123")
	fmt.Println(len("Hello " + "123"))
	fmt.Println(("Hello " + "123")[0])

	str = "Hello,世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}
