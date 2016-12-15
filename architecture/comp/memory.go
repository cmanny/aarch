package comp

const (
  MEM_READ = iota
  MEM_WRITE
  MEM_RESET
)

type MemOp struct {
  Op int
  Addr int
  Data []byte
}

type Memory struct {
  Communicator
  bytes [4096] byte
  opsPerCycle int
}

func (mu *Memory) Init() {
  mu.InitComms()
  mu.opsPerCycle = 2
}

func (mu *Memory) Data() interface{} {
  return ""
}

func (mu *Memory) State() string {
  return ""
}

func (mu *Memory) Cycle() {
  for {
    mu.Recv(CYCLE)

    index := mu.Recv(MEM_IN_1).(int)
    mu.Send(MEM_OUT_1, mu.bytes[index : index+4])
  }
}

func (m *Memory) Fill(bytes []byte, index int) {
  for i := 0; i < len(bytes); i++ {
    m.bytes[i+index] = bytes[i]
  }
}
