package exe

import (
  "github.com/cmanny/aarch/architecture/comp"
  "github.com/cmanny/aarch/architecture/ins"
)

type ArithmeticUnit struct {
  comp.Communicator
  currentIns InsIn
}

func (au *ArithmeticUnit) Init() {
  au.InitComms()
}

func (au *ArithmeticUnit) Data() interface{} {
  return ""
}

func (au *ArithmeticUnit) State() string {
  return ""
}

func (au *ArithmeticUnit) Cycle() {
  in := au.Recv(comp.PIPE_ARITH_IN).(InsIn)
  res := 0
  switch {
  case in.Code == ins.MUL || in.Code == ins.MULI:
    res = in.Op2 * in.Op3
  case in.Code == ins.ADD || in.Code == ins.ADDI:
    res = in.Op2 + in.Op3
  case in.Code == ins.SUB || in.Code == ins.SUBI:
    res = in.Op2 - in.Op3
  case in.Code == ins.XOR || in.Code == ins.XORI:
    res = in.Op2 ^ in.Op3
  }

  au.Send(comp.PIPE_ARITH_OUT, res)

}
