package exe

import (
  "github.com/cmanny/aarch/architecture/ins"
  "github.com/cmanny/aarch/architecture/comp"
)

type ControlUnit struct {
  com* comp.Communicator

  ip int
}

func (cu *ControlUnit) Init() {
  cu.com = &comp.Communicator{}
  cu.com.Init()

  cu.ip = 0

  cu.com.Inputs["ins"] = make(chan interface{}, 1)
  cu.com.Outputs["ip"] = make(chan interface{}, 1)

  cu.com.Outputs["ip"] <- cu.ip
}

func (cu *ControlUnit) Data() interface{} {
  return ""
}

func (cu *ControlUnit) State() string {
  return ""
}

func (cu *ControlUnit) Cycle() {
  bytes := (<-cu.com.Inputs["ins"]).([]byte)

  switch bytes[0] {
  case ins.INS_JMP:
    cu.ip += 4 * int(bytes[1])
  default:
    cu.ip += 4
  }

  cu.com.Outputs["ip"] <- cu.ip
}

func (cu *ControlUnit) Communicator() *comp.Communicator {
  return cu.com
}
