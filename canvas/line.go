package canvas

import "fmt"

func Line(x int, s string) {
  for range x {
    fmt.Print(s)
  }
  fmt.Print("\n")
}
