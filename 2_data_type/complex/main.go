package main

import "fmt"

func main() {
  var val1 complex64 // 2 float32

  val1 = 3.2 + 12i
  val2 := 3.2 + 12i // complex128
  val3 := complex(3.2, 12) // complex128

  fmt.Println(val1)
  fmt.Println(val2)
  fmt.Println(val3)

  fmt.Println(real(val1))
  fmt.Println(imag(val2))
}
