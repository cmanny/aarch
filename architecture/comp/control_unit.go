package comp

import (
  "github.com/cmanny/aarch/architecture/ins"
)

type ControlUnit struct {
  Communicator

}

func (cu *ControlUnit) Init() {
  cu.InitComms()

}

func (cu *ControlUnit) Data() interface{} {
  return ""
}

func (cu *ControlUnit) State() string {
  return ""
}

func (cu *ControlUnit) Cycle() {
  in := cu.Recv(PIPE_CONTROL_IN).(InsIn)
  out := in
  out.Result = out.Ip
  switch {
    case in.Code == ins.JMP:
      out.Result = in.Ip + in.Op1 * 4
    case in.Code == ins.JEQ:
      if in.Op2 == CMP_EQUAL {
        out.Result = in.Ip + in.Op1 * 4
      }
    case in.Code == ins.JNE:
      if in.Op2 != CMP_EQUAL {
        out.Result = in.Ip + in.Op1 * 4
      }
    case in.Code == ins.JL:
      if in.Op2 == CMP_LESS_THAN {
        out.Result = in.Ip + in.Op1 * 4
      }
    case in.Code == ins.JG:
      if in.Op2 != CMP_GREATER_THAN {
        out.Result = in.Ip + in.Op1 * 4
      }
    case in.Code == ins.JRND:
      //lol

    case in.Code == ins.AJMP:
      out.Result = in.Op2
    case in.Code == ins.AJEQ:
      if in.Op2 == CMP_EQUAL {
        out.Result = in.Op2
      }
    case in.Code == ins.AJNE:
      if in.Op2 != CMP_EQUAL {
        out.Result = in.Op2
      }
    case in.Code == ins.AJL:
      if in.Op2 == CMP_LESS_THAN {
        out.Result = in.Op2
      }
    case in.Code == ins.AJG:
      if in.Op2 != CMP_GREATER_THAN {
        out.Result = in.Op2
      }
    case in.Code == ins.HALT:
      out.Result = -1 //Secret IP

    default:
  }

  cu.Send(PIPE_CONTROL_OUT, out)
}
