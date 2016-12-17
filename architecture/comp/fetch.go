package comp

//import "fmt"

import (
  "github.com/cmanny/aarch/architecture/ins"
)

type Fetch struct {
  Communicator
  PipelineData

  is *ins.InstructionSet

  hitControl bool
  ip int
  tag int
  bw int
}

func (fu *Fetch) Init(is *ins.InstructionSet) {
  fu.InitComms()
  fu.ip = 0
  fu.tag = 0
  fu.is = is
  fu.hitControl = false

  fu.current = make([]InsIn, 0)
  fu.bw = 16
}

func (fu *Fetch) Data() interface{} {
  return ""
}

func (fu *Fetch) State() string {
  return ""
}

func (fu *Fetch) Cycle() {
  for {
    fu.Recv(CYCLE)
    fu.Send(PIPE_DECODE_IN, fu.current)
    //fmt.Println("Sending fetched data")

    read := MemOp{}
    read.Op = MEM_READ
    read.Addr = fu.ip
    read.Len = fu.bw
    fu.Send(MEM_IN_1, read)
    bytes := fu.Recv(MEM_OUT_1).([]byte)
    insns := make([]InsIn, 0)

    // Turn bytes into InsIn objects
    if fu.hitControl {
      fu.current = insns
      continue
    }
    for i := 0; i < len(bytes); i += 4 {
      insin := InsIn{}
      insin.Tag = fu.tag
      insin.Ip = fu.ip
      insin.Code = bytes[i]
      insin.RawOp1 = bytes[i + 1]
      insin.RawOp2 = bytes[i + 2]
      insin.RawOp3 = bytes[i + 3]

      insns = append(insns, insin)
      fu.tag++
      fu.ip += 4
      if t, _ := fu.is.InsIdDecode(insin.Code); t.Ins_type == ins.TYPE_CONTROL {
        fu.hitControl = true
        // We hit a control instruction, break
        break
      }
    }
    fu.current = insns

  }
}
