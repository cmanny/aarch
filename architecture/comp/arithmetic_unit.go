package comp

import (
  "github.com/cmanny/aarch/architecture/ins"
  //"fmt"
)

const (
  CMP_EQUAL = iota
  CMP_LESS_THAN
  CMP_GREATER_THAN
)

type ArithmeticUnit struct {
  Communicator
  PipelineData
}

func (au *ArithmeticUnit) Init() {
  au.InitComms()
  au.next = InsIn{}
  au.current = InsIn{}
}

func (au *ArithmeticUnit) Data() interface{} {
  return ""
}

func (au *ArithmeticUnit) State() string {
  return ""
}

func (au *ArithmeticUnit) Cycle() {
  for {
    au.Recv(CYCLE)
    au.current = au.next
    au.Send(PIPE_ARITH_OUT, au.current)
    in := au.Recv(PIPE_ARITH_IN).(InsIn)


    out := in
    switch {
      case in.Code == ins.MUL || in.Code == ins.MULI:
        out.Result = in.Op2 * in.Op3
      case in.Code == ins.ADD || in.Code == ins.ADDI:
        out.Result = in.Op2 + in.Op3
      case in.Code == ins.SUB || in.Code == ins.SUBI:
        out.Result = in.Op2 - in.Op3
      case in.Code == ins.XOR || in.Code == ins.XORI:
        out.Result = in.Op2 ^ in.Op3
      case in.Code == ins.LEAL:
        out.Result = out.Result | (in.Op2<<8|in.Op3)<<0
      case in.Code == ins.LEAH:
        out.Result = out.Result | (in.Op2<<8|in.Op3)<<16
      case in.Code == ins.CMP || in.Code == ins.CMPI:
        if in.Op2 == in.Op3 {
          out.Result = CMP_EQUAL
        }
        if in.Op2 < in.Op3 {
          out.Result = CMP_LESS_THAN
        }
        if in.Op2 > in.Op3 {
          out.Result = CMP_GREATER_THAN
        }
    }
    //fmt.Println(out.Result)
    au.next = out

  }
}
