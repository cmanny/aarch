package architecture

import (
  "fmt"
)

type Processor struct {
  clockSpeed int
  numExUnits int
}

func (p Processor) Run() {
  fmt.Print("Running")
}
