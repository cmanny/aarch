package exe

import (
  "github.com/cmanny/aarch/architecture/comp"
  "github.com/cmanny/aarch/architecture/ins"
)

type ControlUnit struct {
  comp.Communicator

  ip int
}

func (cu *ControlUnit) Init() {
  cu.InitComms()

  cu.ip = 0

  cu.Inputs["ins"] = make(chan interface{}, 1)
  cu.Outputs["ip"] = make(chan interface{}, 1)

  cu.Outputs["ip"] <- cu.ip
}

func (cu *ControlUnit) Data() interface{} {
  return ""
}

func (cu *ControlUnit) State() string {
  return ""
}

func (cu *ControlUnit) Cycle() {
  bytes := (<-cu.Inputs["ins"]).([]byte)

  switch bytes[0] {
  case ins.JMP:
    cu.ip += 4 * int(bytes[1])
  case ins.HALT:
    cu.ip = -1
  default:
    cu.ip += 4
  }

  cu.Outputs["ip"] <- cu.ip
}
