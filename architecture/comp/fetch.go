package comp

type Fetch struct {
	Communicator
}

func (fu *Fetch) Init() {
	fu.InitComms()
	fu.Inputs["mem_in"] = make(chan interface{}, 1)
	fu.Inputs["mem_out"] = make(chan interface{}, 1)
}

func (fu *Fetch) Data() interface{} {
	return ""
}

func (fu *Fetch) State() string {
	return ""
}

func (fu *Fetch) Cycle() {
}
