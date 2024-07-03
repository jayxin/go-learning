package main

import "fmt"

func main() {
	var (
		v1 int
		v2 string
	)
	var v3 [10]int
	var v4 []int
	var v5 struct {
		f int
	}
	var v6 *int
	var v7 map[string]int
	var v8 func(a int) int

	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8)

	var i int = 10
	var j = 1
	l := 10
	i, j = j, i

	fmt.Println(i, j, l)

	const Pi float64 = 3.14
	const zero = 0.0
	const (
		size int64 = 1024
		eof        = -1
	)
	const u, v float32 = 0, 3
	const a, b, c = 3, 4, "foo"

	const mask = 1 << 3

	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)

	const (
		a0 = 1 << iota
		a1
		a2
	)

	const (
		b0         = iota * 42
		b1 float64 = iota * 42
		b2         = iota * 42
	)

	const x = iota
	const y = iota

	const (
		d0 = iota
		d1
		d2
	)

	// enumeration
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
}
