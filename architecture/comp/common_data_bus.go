package comp

type CommonDataBus struct {
  Communicator
}

func (cdb* CommonDataBus) Init() {
  cdb.InitComms()
}

func (cdb* CommonDataBus) Data() interface{} {
  return ""
}

func (cdb* CommonDataBus) State() string {
  return ""
}

func (cdb* CommonDataBus) Cycle() {
  for {
    cdb.Recv(CYCLE)
  }
}
