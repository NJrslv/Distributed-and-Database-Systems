package main

import "fmt"

func fibonacci() func() int {
  f, s := 0, 1
  return func() int {
    res := f
    f, s = s, s + res
    return res
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
