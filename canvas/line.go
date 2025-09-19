package canvas

import "fmt"

func Line(x int, s string, b bool) {
  if b {fmt.Println(s+" ")}
  for range x {
    fmt.Print(s)
  }
  if b {fmt.Print(" "+s)}
  fmt.Print("\n")
}
