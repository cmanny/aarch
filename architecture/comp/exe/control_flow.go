package exe

import (
  "github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type ControlUnit struct {

  Data() interface{}
  State() string
  Cycle()
  Communicator() *Communicator
}
