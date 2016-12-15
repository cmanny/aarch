package comp

import (
  "github.com/cmanny/aarch/architecture/ins"
)

type ControlUnit struct {
  Communicator

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
  bytes := cu.Recv(PIPE_CONTROL_IN).([]byte)

  switch bytes[0] {
  case ins.JMP:
    cu.ip += 4 * int(bytes[1])
  case ins.HALT:
    cu.ip = -1
  default:
    cu.ip += 4
  }

  cu.Send(PIPE_CONTROL_OUT, cu.ip)
}
