package comp

import (
  "github.com/cmanny/aarch/architecture/ins"
  "encoding/binary"
)

type MemoryUnit struct {
  Communicator
}

func (mu *MemoryUnit) Init() {
  mu.InitComms()
}

func (mu *MemoryUnit) Data() interface{} {
  return ""
}

func (mu *MemoryUnit) State() string {
  return ""
}

func (mu *MemoryUnit) Cycle() {
  in := mu.Recv(PIPE_MEMORY_IN).(InsIn)
  out := in
  switch {
    case in.Code == ins.MOV || in.Code == ins.MOVI:
      out.Result = in.Op2
    case in.Code == ins.LDR:
      memOp := MemOp{}
      memOp.Op = MEM_READ
      memOp.Addr = in.Op2
      mu.Send(MEM_IN_2, memOp)
      bytes := mu.Recv(MEM_OUT_2).([]byte)
      res, _ := binary.Varint(bytes)
      out.Result = int(res)
    case in.Code == ins.STR:
      memOp := MemOp{}
      memOp.Op = MEM_WRITE
      memOp.Addr = in.Op1
      memOp.Data = make([]byte, 4)
      binary.PutVarint(memOp.Data, int64(in.Op2))
      mu.Send(MEM_IN_2, memOp)
  }

  mu.Send(PIPE_MEMORY_OUT, out)

}
