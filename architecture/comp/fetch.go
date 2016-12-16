package comp

import "fmt"

type Fetch struct {
  Communicator
  PipelineData

  ip int
}

func (fu *Fetch) Init() {
  fu.InitComms()
  fu.ip = 0
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

    read := MemOp{}
    read.Op = MEM_READ
    read.Addr = fu.ip
    read.Len = 4
    fu.Send(MEM_IN_1, read)
    bytes := fu.Recv(MEM_OUT_1).([]byte)
    for i := 0; i < len(bytes); i += 4 {
      fmt.Println(bytes)
    }
    fu.ip += 4

  }
}
