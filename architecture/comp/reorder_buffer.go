package comp

import (
  //"fmt"
  "container/list"
  "github.com/cmanny/aarch/architecture/ins"
)

const (
  RB_ISSUED = iota
  RB_EXECUTING
  RB_FINISHED
)

type BufferEntry struct {
  in InsIn
  state int
}

type ReorderBuffer struct {
  Communicator
  is *ins.InstructionSet

  buffer *list.List
  decoded *list.List
}

func (rb* ReorderBuffer) Init(is *ins.InstructionSet) {
  rb.InitComms()
  rb.buffer = list.New()
  rb.decoded = list.New()
  rb.is = is
}

func (rb* ReorderBuffer) Data() interface{} {
  return ""
}

func (rb* ReorderBuffer) State() string {
  return ""
}

func (rb* ReorderBuffer) Cycle() {
  for {
    rb.Recv(CYCLE)
    rb.Send(PIPE_RS_IN, rb.decoded)
    decoded := rb.Recv(PIPE_DECODE_OUT).([]InsIn)

    entryList := list.New()
    rsList := list.New()

    for _, in := range decoded {
      entry := BufferEntry{}
      entry.state = RB_ISSUED
      entry.in = in

      entryList.PushBack(entry)
      rsList.PushBack(in)
    }
    rb.buffer.PushBackList(entryList)
    rb.decoded = rsList
  }
}
