package main

import (
  "fmt"

  "github.com/kyokomi/emoji"
)

func main() {
  i := 0
  sushi := emoji.Sprint(":sushi:")
  for {
    fmt.Println(sushi)
    i++
    if i == 10 {
      break
    }
  }
}
