package comp

type Decode struct {
  Communicator
}

func (fu *Decode) Init() {
  fu.InitComms()
  fu.Inputs["reg_in"] = make(chan interface{}, 1)
  fu.Outputs["reg_out"] = make(chan interface{}, 1)
}

func (fu *Decode) Data() interface{} {
  return ""
}

func (fu *Decode) State() string {
  return ""
}

func (fu *Decode) Cycle() {
}
