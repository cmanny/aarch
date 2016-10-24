package exe

import (
  "github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type ControlUnit struct {
  com* Communicator
}

func (cu *ControlUnit) Init() {
  cu.com = &Communicator{}
}

func (cu *ControlUnit) Data() interface{} {
  return ""
}

func (cu *ControlUnit) State() string {
  return ""
}

func (cu *ControlUnit) Cycle() {

}

func (cu *ControlUnit) Communicator() *Communicator {
  return cu.com
}
