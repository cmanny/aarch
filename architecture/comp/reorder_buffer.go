package comp

type ReorderBuffer struct {
  Communicator
}

func (rb* ReorderBuffer) Init() {
  rb.InitComms()
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
