package comp

import "fmt"


type Decode struct {
  Communicator
  PipelineData
}

func (du *Decode) Init() {
  du.InitComms()
  du.current = make([]InsIn, 0)
  du.next = make([]InsIn, 0)
}

func (du *Decode) Data() interface{} {
  return ""
}

func (du *Decode) State() string {
  return ""
}

func (du *Decode) Cycle() {
  for {
    du.Recv(CYCLE)
    du.current = du.next
    du.Send(PIPE_RS_IN, du.current)
    du.next = du.Recv(PIPE_DECODE_IN)

    // Decode instructions
    for _, insn := range du.current.([]InsIn) {
      fmt.Println(insn)
    }

  }
}
