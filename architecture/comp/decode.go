package comp

type Decode struct {
  Communicator
  PipelineData
}

func (fu *Decode) Init() {
  fu.InitComms()
}

func (fu *Decode) Data() interface{} {
  return ""
}

func (fu *Decode) State() string {
  return ""
}

func (fu *Decode) Cycle() {
  for {
    fu.Recv(CYCLE)
  }
}
