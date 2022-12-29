package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(buf io.Writer) {
  for i := 3; i > 0; i-- {
    fmt.Fprintln(buf, i)
  }
  fmt.Fprint(buf, "Go!")
}

func main() {
  Countdown(os.Stdout)
}

