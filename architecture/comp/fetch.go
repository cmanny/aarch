package comp

type Fetch struct {
  Communicator
}

func (fu *Fetch) Init() {
  fu.InitComms()
}

func (fu *Fetch) Data() interface{} {
  return ""
}

func (fu *Fetch) State() string {
  return ""
}

func (fu *Fetch) Cycle() {
}
