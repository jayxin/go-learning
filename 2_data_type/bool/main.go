package main

import "fmt"

func main() {
  var v1 bool
  v1 = true

  //v1 = 1 // error!
  //v1 = bool(1) // error!

  fmt.Println(v1)

  v2 := (1 == 2)

  fmt.Println(v2)
}
