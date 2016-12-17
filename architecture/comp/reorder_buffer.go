package comp

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

  buffer []BufferEntry
  decoded []InsIn
}

func (rb* ReorderBuffer) Init() {
  rb.InitComms()
  rb.buffer = make([]BufferEntry, 0)
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

    new := make([]BufferEntry, len(rb.decoded))
    for i, in := range rb.decoded {
      new[i] = BufferEntry{}
      new[i].state = RB_ISSUED
      new[i].in = in
    }
    rb.buffer = append(rb.buffer, new...)
    rb.decoded = rb.Recv(PIPE_DECODE_OUT).([]InsIn)
  }
}
