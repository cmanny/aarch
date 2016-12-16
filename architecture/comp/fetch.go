package comp

type Fetch struct {
  Communicator
  PipelineData

  ip int
  bw int
}

func (fu *Fetch) Init() {
  fu.InitComms()
  fu.ip = 0
  fu.current = make([]InsIn, 0)
  fu.bw = 4
}

func (fu *Fetch) Data() interface{} {
  return ""
}

func (fu *Fetch) State() string {
  return ""
}

func (fu *Fetch) Cycle() {
  for {
    fu.Recv(CYCLE)
    fu.Send(PIPE_DECODE_IN, fu.current)

    read := MemOp{}
    read.Op = MEM_READ
    read.Addr = fu.ip
    read.Len = fu.bw
    fu.Send(MEM_IN_1, read)
    bytes := fu.Recv(MEM_OUT_1).([]byte)
    insns := make([]InsIn, fu.bw / 4)

    // Turn bytes into InsIn objects
    for i := 0; i < len(bytes); i += 4 {
      insns[i] = InsIn{}
      insns[i].Ip = fu.ip + i * 4
      insns[i].Code = bytes[i]
      insns[i].RawOp1 = bytes[i + 1]
      insns[i].RawOp2 = bytes[i + 2]
      insns[i].RawOp3 = bytes[i + 3]
    }
    fu.current = insns
    fu.ip += fu.bw

  }
}
