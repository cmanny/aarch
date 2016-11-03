package comp

type Decode struct {
  Communicator
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
}
