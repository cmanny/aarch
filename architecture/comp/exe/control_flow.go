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

}

func (cu *ControlUnit) Data() interface{} {
  return ""
}

func (cu *ControlUnit) State() string {
  return ""
}

func (cu *ControlUnit) Cycle() {
  bytes := cu.Recv(comp.PIPE_CONTROL_IN).([]byte)

  switch bytes[0] {
  case ins.JMP:
    cu.ip += 4 * int(bytes[1])
  case ins.HALT:
    cu.ip = -1
  default:
    cu.ip += 4
  }

  cu.Send(comp.PIPE_CONTROL_OUT, cu.ip)
}
