package main

import "errors"
import "fmt"

func main() {
	res, ok := Add(1, 2)
	if ok != nil {
		fmt.Println("Error!")
		return
	} else {
		fmt.Println(res)
	}

	myfunc(1, 23, 4)
	myfunc(1, 2)

	myfunc1([]int{1, 23, 4})
	myfunc1([]int{1, 2})

	var v1 int = 1
	var v2 int64 = 234
	var v3 string = "hello"
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)

	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(1, 2))
	/*func(ch chan int){
	  ch<-ACK
	}(reply_chan)*/

  // closure
	var j int = 5
	a := func() (func()) {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
  a()
  j *= 2
  a()
}

// func Add(a int, b int) (ret int, err error) {
func Add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("Should be non-negative integers!")
		return
	}
	return a + b, nil
}

func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func myfunc1(args []int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func myfunc2(args ...int) {
	myfunc(args...)
	myfunc(args[1:]...)
}

/*func Printf(format string, args ...interface{}) {
  // ...
}*/

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, " is an int value.")
		case string:
			fmt.Println(arg, " is an string value.")
		case int64:
			fmt.Println(arg, " is an int64 value.")
		default:
			fmt.Println(arg, " is an unknown type.")
		}
	}
}
