package comp

const (
  MEM_READ = iota
  MEM_WRITE
  MEM_RESET
  MEM_CHAN1_IN
  MEM_CHAN2_IN
  MEM_CHAN3_IN
  MEM_CHAN1_OUT
  MEM_CHAN2_OUT
  MEM_CHAN3_OUT
)

type MemOp struct {
  op int
  addr int
  data []byte
}

type Memory struct {
  Communicator
  bytes [4096]byte
  opsPerCycle int
}

func (mu *Memory) Init() {
  mu.InitComms()
  mu.opsPerCycle = 2
  mu.Inputs["in1"] = make(chan interface{}, 1)
  mu.Outputs["out1"] = make(chan interface{}, 1)
}

func (mu *Memory) Data() interface{} {
  return ""
}

func (mu *Memory) State() string {
  return ""
}

func (mu *Memory) Cycle() {
  index := (<-mu.Inputs["in1"]).(int)
  mu.Outputs["out1"] <- mu.bytes[index : index+4]
}

func (m *Memory) Fill(bytes []byte, index int) {
  for i := 0; i < len(bytes); i++ {
    m.bytes[i+index] = bytes[i]
  }
}
