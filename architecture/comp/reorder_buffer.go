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
}

func (rb* ReorderBuffer) Init() {
  rb.InitComms()
  rb.buffer = make([]BufferEntry, 0)
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


  }
}
