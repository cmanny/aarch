package main

import (
  //"fmt"

  "github.com/cmanny/aarch/architecture"
)

func main() {
  p := &architecture.Processor{}
  p.Debug(true)
  p.Run()
}
