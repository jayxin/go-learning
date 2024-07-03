package main

import "fmt"

func main() {
  value1 := 64 // int

  //var value2 int32
  //value2 = value1 // error

  var value3 int32
  value3 = int32(value1)

  fmt.Println(value3)

  fmt.Println(1 + 1)
  fmt.Println(1 - 1)
  fmt.Println(1 * 1)
  fmt.Println(1 / 1)
  fmt.Println(1 % 1)

  fmt.Println(1 > 1)
  fmt.Println(1 < 1)
  fmt.Println(1 >= 1)
  fmt.Println(1 <= 1)
  fmt.Println(1 != 1)
  fmt.Println(1 == 1)

  i, j := 1, 2
  if i == j {
    fmt.Println("i and  are equal.")
  }

  var k int32
  var l int64

  k , l = 1, 2

  /*if k == l { // error!
    fmt.Println("k and l are equal.")
  }*/

  if k == 1 || l == 2 {
    fmt.Println("Pass!")
  }

  fmt.Println(124 << 2)
  fmt.Println(124 >> 2)
  fmt.Println(124 ^ 2)
  fmt.Println(124 & 2)
  fmt.Println(124 | 2)
  fmt.Println(^2) // bit-wise not
}
