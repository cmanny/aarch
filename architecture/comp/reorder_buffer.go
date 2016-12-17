package comp

import (
  "fmt"
  "container/list"
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

  buffer *list.List
  decoded []InsIn
}

func (rb* ReorderBuffer) Init() {
  rb.InitComms()
  rb.buffer = list.New()
  rb.decoded = make([]InsIn, 0)
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
    rb.decoded = rb.Recv(PIPE_DECODE_OUT).([]InsIn)

    for _, in := range rb.decoded {
      entry := BufferEntry{}
      entry.state = RB_ISSUED
      entry.in = in

      rb.buffer.PushBack(entry)
    }
    fmt.Println(rb.buffer)
  }
}
